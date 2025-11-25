package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	vanityPattern := os.Getenv("VANITY_PATTERN")
	if vanityPattern == "" {
		vanityPattern = "on1y"
	}

	targetCount, _ := strconv.Atoi(os.Getenv("TARGET_KEYPAIR_COUNT"))
	if targetCount == 0 {
		targetCount = 50
	}

	numWorkers, _ := strconv.Atoi(os.Getenv("NUM_WORKERS"))
	if numWorkers == 0 {
		numWorkers = 8
	}

	fmt.Println("Vanity Generator Worker Started (Go)")
	fmt.Printf("Target: %d keypairs ending with \"%s\"\n", targetCount, vanityPattern)
	fmt.Printf("Workers: %d concurrent goroutines\n", numWorkers)
	fmt.Println(strings.Repeat("━", 60))

	db, err := NewDatabase(dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	currentCount, err := db.GetAvailableCount()
	if err != nil {
		log.Fatalf("Failed to get available count: %v", err)
	}

	fmt.Printf("Current available keypairs: %d\n", currentCount)

	if currentCount >= targetCount {
		fmt.Printf("Target already met! (%d/%d)\n", currentCount, targetCount)
		return
	}

	needed := targetCount - currentCount
	fmt.Printf("Need to generate %d more keypairs\n", needed)
	fmt.Println(strings.Repeat("━", 60))

	startTime := time.Now()
	var completed atomic.Int32
	var totalAttempts atomic.Int64

	resultChan := make(chan VanityResult, numWorkers)
	var wg sync.WaitGroup

	workerQueue := make(chan int, needed)
	for i := 0; i < needed; i++ {
		workerQueue <- i
	}
	close(workerQueue)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for range workerQueue {
				GenerateVanityKeypair(vanityPattern, resultChan, workerID)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		totalAttempts.Add(int64(result.Attempts))

		fmt.Printf("\nFound vanity address after %d attempts: %s\n", result.Attempts, result.PublicKey)

		exists, err := db.KeypairExists(result.PublicKey)
		if err != nil {
			log.Printf("Error checking keypair existence: %v", err)
			continue
		}
		if exists {
			fmt.Println("Keypair already exists, skipping...")
			continue
		}

		secretKeyBase58 := result.Keypair.String()
		secretKeyBytes := []byte(result.Keypair)

		err = db.CreateKeypair(result.PublicKey, secretKeyBase58, secretKeyBytes, vanityPattern, result.Attempts)
		if err != nil {
			log.Printf("Error saving to database: %v", err)
			continue
		}
		fmt.Printf("Saved to database: %s\n", result.PublicKey)

		completed.Add(1)
		current := completed.Load()
		fmt.Printf("Progress: %d/%d completed\n", current, needed)
	}

	duration := time.Since(startTime)
	fmt.Println("\n" + strings.Repeat("━", 60))
	fmt.Println("All keypairs generated successfully!")

	available, total, avgAttempts, err := db.GetStats()
	if err != nil {
		log.Printf("Error getting stats: %v", err)
	} else {
		fmt.Println("\nFinal Statistics:")
		fmt.Printf("   - Available: %d\n", available)
		fmt.Printf("   - Total: %d\n", total)
		fmt.Printf("   - Avg Attempts: %.0f\n", avgAttempts)
		fmt.Printf("   - Total Time: %s\n", duration.Round(time.Second))
		fmt.Printf("   - Keypairs/sec: %.2f\n", float64(needed)/duration.Seconds())
		fmt.Printf("   - Total Attempts: %d\n", totalAttempts.Load())
	}
}

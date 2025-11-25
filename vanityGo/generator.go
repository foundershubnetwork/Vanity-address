package main

import (
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"
)

type VanityResult struct {
	Keypair   solana.PrivateKey
	PublicKey string
	Attempts  int
}

func GenerateVanityKeypair(pattern string, resultChan chan<- VanityResult, workerID int) {
	attempts := 0

	for {
		attempts++

		account := solana.NewWallet()
		publicKey := account.PublicKey().String()

		if strings.HasSuffix(publicKey, strings.ToLower(pattern)) {
			resultChan <- VanityResult{
				Keypair:   account.PrivateKey,
				PublicKey: publicKey,
				Attempts:  attempts,
			}
			return
		}

		if attempts%10000 == 0 {
			fmt.Printf("[Worker %d] Tried %d addresses...\n", workerID, attempts)
		}
	}
}

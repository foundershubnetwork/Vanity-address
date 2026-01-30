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

func GenerateVanityKeypair(prefix, suffix string, resultChan chan<- VanityResult, workerID int) {
	attempts := 0

	for {
		attempts++

		account := solana.NewWallet()
		publicKey := account.PublicKey().String()

		// Case-sensitive matching for now, as is standard for vanity addresses
		// User can handle casing in the input string if needed
		if strings.HasPrefix(publicKey, prefix) && strings.HasSuffix(publicKey, suffix) {
			resultChan <- VanityResult{
				Keypair:   account.PrivateKey,
				PublicKey: publicKey,
				Attempts:  attempts,
			}
			return
		}

		if attempts%100000 == 0 {
			fmt.Printf("[Worker %d] Tried %d addresses...\n", workerID, attempts)
		}
	}
}

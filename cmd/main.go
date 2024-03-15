package main

import (
	blockchain "github.com/arturbaldoramos/go-crypto/pkg"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddTransaction("A", "B", 1.0)
	previousHash := bc.LastBlock().Hash()
	nonce := bc.ProofOfWork()
	bc.CreateBlock(nonce, previousHash)

	bc.AddTransaction("C", "D", 2.0)
	bc.AddTransaction("X", "Y", 3.0)
	previousHash = bc.LastBlock().Hash()
	nonce = bc.ProofOfWork()
	bc.CreateBlock(nonce, previousHash)

	bc.Print()
}

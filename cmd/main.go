package main

import (
	blockchain "github.com/arturbaldoramos/go-crypto/pkg/block"
)

func main() {
	myBlockchainAddress := "minha_blockchain"
	bc := blockchain.NewBlockchain(myBlockchainAddress)

	bc.AddTransaction("A", "B", 1.0)
	bc.Mining()

	bc.AddTransaction("C", "D", 2.0)
	bc.AddTransaction("X", "Y", 3.0)
	bc.Mining()

	bc.Print()
}

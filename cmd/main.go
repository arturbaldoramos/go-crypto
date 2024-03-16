package main

import (
	"fmt"
	"github.com/arturbaldoramos/go-crypto/pkg/wallet"
)

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
}

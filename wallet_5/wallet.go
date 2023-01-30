package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"fmt"
    "transactions"
)

type Wallet struct {
	PrivateKey []byte
	PublicKey  []byte
}

func (w *Wallet) SignTransaction(transaction Transaction) []byte {
	r, s, err := elliptic.Sign(rand.Reader, w.PrivateKey, sha256.Sum256(transaction))
	if err != nil {
		log.Fatal(err)
	}
	return append(r, s...)
}

func NewWallet() *Wallet {
	privateKey, err := elliptic.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := elliptic.Marshal(privateKey.Curve, privateKey.X, privateKey.Y)
	return &Wallet{privateKey, publicKey}
}

func main() {
	wallet := NewWallet()
	transaction := Transaction{Sender: "A", Recipient: "B", Amount: 1}
	signature := wallet.SignTransaction(transaction)
	fmt.Printf("Signature: %x\n", signature)
}

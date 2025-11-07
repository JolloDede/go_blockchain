package pkg

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
)

// CreateWallet creates and returns a pointer to a new Wallet initialized for use.
func CreateWallet(name string) *Wallet {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		println("Error Wallet")
	}

	return &Wallet{name: name, privateKey: privateKey, PublicKey: &privateKey.PublicKey}
}

// Wallet represents a simple wallet model.
//
// Fields:
// - Name: human-friendly name for the wallet.
// - privateKey: a reference to the private key of the wallet
// - PublicKey: a public facing key that allows you to send money to a adress
type Wallet struct {
	name       string
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// MakeTransaction returns a singed transaction
func (w *Wallet) MakeTransaction(recieverPK *rsa.PublicKey, amount float64) *Transaction {
	t := CreateTransaction(w.PublicKey, recieverPK, amount)
	t.SetSignature(w.signTransaction(t, recieverPK))
	return t
}

// signTransaction hashes the transactions and returns the singed hash
func (w *Wallet) signTransaction(t *Transaction, recieverPublicKey *rsa.PublicKey) string {
	transJson, err := json.Marshal(t)

	if err != nil {
		println("Failed to marschal transaction to json")
	}

	hash := sha256.Sum256(transJson)

	signature, err := rsa.EncryptPKCS1v15(rand.Reader, recieverPublicKey, hash[:])

	if err != nil {
		println("Signing the transaction failed")
	}

	return string(signature)
}

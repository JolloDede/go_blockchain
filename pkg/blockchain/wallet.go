package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
)

// CreateWallet creates and returns a pointer to a new Wallet initialized for use.
func createWallet() *Wallet {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		println("Error Wallet")
	}

	return &Wallet{privateKey: privateKey, PublicKey: &privateKey.PublicKey}
}

// Wallet represents a simple wallet model.
// / Fields:
// - privateKey: a reference to the private key of the wallet
// - PublicKey: a public facing key that allows you to send money to a adress
type Wallet struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// MakeTransaction returns a singed transaction
func (w *Wallet) MakeTransaction(publicKey *rsa.PublicKey, amount float64) (*Transaction, error) {
	t := CreateTransaction(w.PublicKey, publicKey, amount)
	signature, err := w.signTransaction(t)

	if err != nil {
		println("Failed to sign transaction")
		return nil, err
	}
	t.SetSignature(signature)

	return t, nil
}

// signTransaction signs the transaction with the wallet's private key
func (w *Wallet) signTransaction(t *Transaction) (string, error) {
	transJson, err := json.Marshal(t)

	if err != nil {
		println("Failed to marschal transaction to json")
		return "", err
	}

	hash := sha256.Sum256(transJson)

	signature, err := rsa.SignPKCS1v15(rand.Reader, w.privateKey, crypto.SHA256, hash[:])

	if err != nil {
		println("Signing the transaction failed")
	}

	return string(signature), nil
}

func (w *Wallet) VerifyTransaction(transaction *Transaction, recieverKey *rsa.PublicKey) error {
	transJson, err := json.Marshal(transaction)

	if err != nil {
		println("Failed to marschal transaction to json")
		return err
	}

	hash := sha256.Sum256(transJson)

	err = rsa.VerifyPKCS1v15(recieverKey, crypto.SHA256, hash[:], []byte(transaction.signature))

	return nil
}

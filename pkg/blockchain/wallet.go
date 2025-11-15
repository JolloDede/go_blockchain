package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"sync"
)

// CreateWallet creates and returns a pointer to a new Wallet initialized for use.
func createWallet() *Wallet {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		println("Error Wallet")
	}

	return &Wallet{privateKey: privateKey, balance: 0}
}

// Wallet represents a simple wallet model.
// / Fields:
// - privateKey: a reference to the private key of the wallet
// - PublicKey: a public facing key that allows you to send money to a adress
// - balance: the balance of the wallet
// - mu: a mutex for synchronizing access to the balance
type Wallet struct {
	privateKey *rsa.PrivateKey
	balance    float64
	mu         sync.Mutex
}

// MakeTransaction returns a singed transaction
func (w *Wallet) MakeTransaction(publicKey *rsa.PublicKey, amount float64) (*Transaction, error) {
	t := CreateTransaction(w.GetPublicKey(), publicKey, amount)
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

// GetBalance returns the balance of the wallet
func (w *Wallet) GetBalance() float64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.balance
}

// setBalance sets the balance of the wallet
//
// This is only to be used by the blockchain when updating the balance after a block is mined
func (w *Wallet) setBalance(amount float64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance = amount
}

// GetPublicKey returns the public key of the wallet
func (w *Wallet) GetPublicKey() *rsa.PublicKey {
	return &w.privateKey.PublicKey
}

// VerifyTransaction verifies the transaction with the reciever's public key
func VerifyTransaction(transaction *Transaction, recieverKey *rsa.PublicKey) error {
	transJson, err := json.Marshal(transaction)

	if err != nil {
		println("Failed to marschal transaction to json")
		return err
	}

	hash := sha256.Sum256(transJson)

	err = rsa.VerifyPKCS1v15(recieverKey, crypto.SHA256, hash[:], []byte(transaction.signature))

	return nil
}

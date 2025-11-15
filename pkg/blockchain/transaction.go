package blockchain

import (
	"crypto/rsa"
	"errors"
)

// Function to have a unified interface for creating transactions
func CreateTransaction(sender *rsa.PublicKey, reciever *rsa.PublicKey, amount float64) *Transaction {
	return &Transaction{Sender: sender, Reciever: reciever, Amount: amount}
}

// Transaction is the type of data that we whant to store in our chain.
type Transaction struct {
	Sender    *rsa.PublicKey
	Reciever  *rsa.PublicKey
	Amount    float64
	signature string
}

// Set the signate of a transaction
//
// If the signature is already set then a error occurs because you are only allow to set the signature once
func (t *Transaction) SetSignature(signature string) error {
	if len(t.signature) == 0 {
		return errors.New("Cant set the signature after it is set")
	}

	t.signature = signature
	return nil
}

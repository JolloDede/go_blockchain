package user

import (
	"crypto/rsa"
	"errors"

	"github.com/JolloDede/go_blockchain/pkg/blockchain"
)

// NewUser creates and returns a pointer to a new User initialized for use.
func NewUser(name string, desc string, blockchain *blockchain.Blockchain) *User {
	return &User{name: name, description: desc, blockchain: blockchain, friends: make([]*rsa.PublicKey, 0)}
}

// User represents a simple user model.
// / Fields:
// - name: the name of the user
// - description: a brief description of the user
// - wallet: a reference to the user's wallet
// - blockchain: a reference to the blockchain the user interacts with
// - friends: a slice of public keys representing the user's friends
type User struct {
	name        string
	description string
	wallet      *blockchain.Wallet
	blockchain  *blockchain.Blockchain
	friends     []*rsa.PublicKey
}

// AddWallet assigns a wallet to the user
func (u *User) AddWallet(wallet *blockchain.Wallet) error {
	if u.wallet != nil {
		return errors.New("user already has a wallet")
	}
	u.wallet = wallet

	return nil
}

// MineBlock mines a new block and adds it to the blockchain
func (u *User) MineBlock() {
	block := blockchain.CreateBlock(u.blockchain.GetLastBlock().Hash, u.blockchain.GetPendingTransactions())
	difficulty := int32(0)
	for {
		difficulty++
		block.ProofOfWork(difficulty)
		err := u.blockchain.AddBlock(block)

		if errors.Is(err, blockchain.NewErrBlockAlreadyMinded()) {
			u.MineBlock()
			return
		}
		if err == nil {
			break
		}
	}
}

// MakeTransaction creates a new transaction from the user's wallet
func (u *User) MakeTransaction(reciever *rsa.PublicKey, amount float64) error {
	t, err := u.wallet.MakeTransaction(reciever, amount)

	if err != nil {
		return err
	}

	err = u.blockchain.AddTransaction(t)

	if err != nil {
		return err
	}

	return nil
}

// AddFriend adds a friend's public key to the user's friend list
func (u *User) AddFriend(friend *rsa.PublicKey) {
	u.friends = append(u.friends, friend)
}

// GetFriend returns the public key of a friend at the specified index
func (u *User) GetFriend(index int) *rsa.PublicKey {
	return u.friends[index]
}

// GivePublicKey returns the user's public key
func (u *User) GivePublicKey() *rsa.PublicKey {
	return u.wallet.GetPublicKey()
}

// GetWallet returns the user's wallet
func (u *User) GetWallet() *blockchain.Wallet {
	return u.wallet
}

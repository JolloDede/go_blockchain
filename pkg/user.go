package pkg

import "crypto/rsa"

func NewUser(name string, desc string, blockchain *Blockchain) *User {
	return &User{name: name, description: desc, blockchain: blockchain}
}

type User struct {
	name        string
	description string
	wallet      *Wallet
	blockchain  *Blockchain
}

func (u *User) AddWallet(wallet *Wallet) *rsa.PublicKey {
	u.wallet = wallet
	return u.wallet.PublicKey
}

func (u *User) MineBlock(transactions []*Transaction) {
	block := CreateBlock(u.blockchain.Chain[len(u.blockchain.Chain)-1].Hash, transactions)
	u.blockchain.AddBlock(block)
}

func (u *User) MakeTransaction(reciver *rsa.PublicKey, amount float64) *Transaction {
	return u.wallet.MakeTransaction(reciver, amount)
}

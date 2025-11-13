package blockchain

import (
	"errors"
	"sync"
)

// Function that creates a new blockchain with the genesis block
func CreateBlockchain() *Blockchain {
	bc := &Blockchain{Chain: []*Block{CreateGenesisBlock()}, transactions: make([]*Transaction, 0), Wallets: make([]*Wallet, 0)}
	return bc
}

// Blockchain struct that holds all of the blocks in a slice
//
// Blockchain is the main structure that holds all of the blocks in the chain
// It manages transactions and validates new blocks before adding them to the chain
type Blockchain struct {
	transactions []*Transaction
	Chain        []*Block
	Wallets      []*Wallet
	mu           sync.Mutex
}

// AddTransaction adds a new transaction to the list of pending transactions
func (bc *Blockchain) AddBlock(b *Block) error {
	const difficulty = 4
	if b.ValidateDifficulty(difficulty) && b.PrevHash == bc.GetLastBlock().Hash {
		bc.mu.Lock()
		bc.transactions = bc.transactions[len(b.transactions):]
		bc.Chain = append(bc.Chain, b)
		bc.mu.Unlock()
		return nil
	} else {
		return errors.New("invalid block")
	}
}

// GetPendingTransactions returns the first 10 pending transactions
func (bc *Blockchain) GetPendingTransactions() []*Transaction {
	remove := min(len(bc.transactions), 10)
	return bc.transactions[:remove]
}

// GetLastBlock returns the last block in the blockchain
func (bc *Blockchain) GetLastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

// AddWallet creates a new wallet and adds it to the blockchain's wallet list
func (bc *Blockchain) AddWallet() *Wallet {
	w := createWallet()
	bc.mu.Lock()
	bc.Wallets = append(bc.Wallets, w)
	bc.mu.Unlock()
	return w
}

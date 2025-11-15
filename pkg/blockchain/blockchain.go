package blockchain

import (
	"crypto/x509"
	"errors"
	"sync"
)

// Function that creates a new blockchain with the genesis block
func CreateBlockchain() *Blockchain {
	bc := &Blockchain{Chain: []*Block{CreateGenesisBlock()}, transactions: make([]*Transaction, 0), Wallets: sync.Map{}}
	return bc
}

// Blockchain struct that holds all of the blocks in a slice
//
// Blockchain is the main structure that holds all of the blocks in the chain
// It manages transactions and validates new blocks before adding them to the chain
// / Fields:
// // - transactions: a slice of pending transactions to be included in new blocks
// // - Chain: a slice of blocks representing the blockchain
// // - Wallets: a map of public keys to wallets for managing user balances
// - mu: a mutex for synchronizing access to the blockchain
type Blockchain struct {
	transactions []*Transaction
	Chain        []*Block
	Wallets      sync.Map // map[string(x509.MarshalPKCS1PublicKey(w.PublicKey))]*Wallet
	mu           sync.Mutex
}

// AddTransaction adds a new transaction to the list of pending transactions
func (bc *Blockchain) AddBlock(b *Block) error {
	const difficulty = 4

	if !b.ValidateDifficulty(difficulty) {
		return errors.New("Block isn't difficult enough")
	}
	if b.PrevHash != bc.GetLastBlock().Hash {
		return NewErrBlockAlreadyMinded()
	}
	bc.mu.Lock()
	bc.transactions = bc.transactions[len(b.transactions):]
	bc.Chain = append(bc.Chain, b)
	bc.mu.Unlock()

	return nil
}

// GetPendingTransactions returns the first 10 pending transactions
func (bc *Blockchain) GetPendingTransactions() []*Transaction {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	remove := min(len(bc.transactions), 10)
	return bc.transactions[:remove]
}

// GetLastBlock returns the last block in the blockchain
func (bc *Blockchain) GetLastBlock() *Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	return bc.Chain[len(bc.Chain)-1]
}

// AddWallet creates a new wallet and adds it to the blockchain's wallet list
func (bc *Blockchain) AddWallet() *Wallet {
	w := createWallet()

	bc.mu.Lock()
	bc.Wallets.Store(string(x509.MarshalPKCS1PublicKey(w.PublicKey)), w)
	bc.mu.Unlock()

	return w
}

// AddTransaction adds a new transaction to the list of pending transactions
// It also updates the balances of the sender and receiver wallets
func (bc *Blockchain) AddTransaction(t *Transaction) error {
	err := VerifyTransaction(t, t.Reciever)

	if err != nil {
		return err
	}

	val, ok := bc.Wallets.Load(string(x509.MarshalPKCS1PublicKey(t.Sender)))
	if !ok {
		return errors.New("Sender wallet not found")
	}
	wallet := val.(*Wallet)
	wallet.setBalance(wallet.GetBalance() - t.Amount)
	val2, ok := bc.Wallets.Load(string(x509.MarshalPKCS1PublicKey(t.Reciever)))
	if !ok {
		return errors.New("Receiver wallet not found")
	}
	recieverWallet := val2.(*Wallet)
	recieverWallet.setBalance(recieverWallet.GetBalance() + t.Amount)

	bc.mu.Lock()
	bc.transactions = append(bc.transactions, t)
	bc.mu.Unlock()

	return nil
}

// ErrBlockAlreadyMinded is an error indicating that a block has already been mined
type ErrBlockAlreadyMinded struct{}

// NewErrBlockAlreadyMinded creates a new ErrBlockAlreadyMinded error
func NewErrBlockAlreadyMinded() error {
	return &ErrBlockAlreadyMinded{}
}

// Error returns the error message for ErrBlockAlreadyMinded
func (e *ErrBlockAlreadyMinded) Error() string {
	return "block has already been mined"
}

package pkg

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// Creates the first Block in the Blockchain called Genesis
func CreateGenesisBlock() *Block {
	transaction := &Transaction{Sender: "Genesis", Reciever: "Genesis", Amount: 0.0}
	return CreateBlock("", []*Transaction{transaction})
}

// Function for creating all of the blocks in the chain
func CreateBlock(prevHash string, transactions []*Transaction) *Block {
	t := time.Now()
	stamp := t.Format(time.Stamp)
	b := &Block{timestamp: stamp, PrevHash: prevHash, transactions: transactions, Nonce: 0}
	b.Hash = b.CalculateHash()
	return b
}

// The main structure the blockchain revolves around.
// The transaction, timestamp and staticData are private
// because they shouldnt change after the creation.
type Block struct {
	timestamp    string
	Hash         string
	PrevHash     string
	transactions []*Transaction
	Nonce        int32
	staticData   string
}

// This function calculates the hash of the block and returns it
func (b *Block) CalculateHash() string {
	if b.staticData == "" {
		transJson, err := json.Marshal(b.transactions)

		if err != nil {
			println("lul")
		}

		b.staticData = string(transJson)
	}

	data := fmt.Sprintf("%d%s", b.Nonce, b.staticData)

	hash := sha256.Sum256([]byte(data))

	return fmt.Sprintf("%x", hash)
}

// In this function the CalculateHash is called until the the difficulty of the hash is enough
func (b *Block) ProofOfWork(difficulty int32) {
	var difficultyString string
	for i := 0; i < int(difficulty); i++ {
		difficultyString += string('0')
	}
	for {
		b.Nonce++
		b.Hash = b.CalculateHash()
		if b.Hash[:difficulty] == difficultyString {
			println("Block mined: ", b.Hash)
			break
		}
	}
}

package pkg

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

func CreateGenesisBlock() *Block {
	transaction := &Transaction{Sender: "Genesis", Reciever: "Genesis", Amount: 0.0}
	return CreateBlock("", []*Transaction{transaction})
}

func CreateBlock(prevHash string, transactions []*Transaction) *Block {
	t := time.Now()
	stamp := t.Format(time.Stamp)
	b := &Block{Timestamp: stamp, PrevHash: prevHash, Transactions: transactions, Nonce: 0}
	b.Hash = b.CalculateHash()
	return b
}

type Block struct {
	Timestamp    string
	Hash         string
	PrevHash     string
	Transactions []*Transaction
	Nonce        int32
}

func (b *Block) CalculateHash() string {
	transJson, err := json.Marshal(b.Transactions)

	if err != nil {
		println("lul")
	}

	data := fmt.Sprintf("%d%s%s%s", b.Nonce, b.Timestamp, b.PrevHash, transJson)

	hash := sha256.Sum256([]byte(data))

	return fmt.Sprintf("%x", hash)
}

func (b *Block) ProofOfWork(difficulty int32) {
	var difficultyString string
	for i := 0; i < int(difficulty+1); i++ {
		difficultyString += string('0')
	}
	println(difficultyString)
	for {
		b.Nonce++
		b.Hash = b.CalculateHash()
		if b.Hash[:difficulty+1] == difficultyString {
			println("Block minded: ", b.Hash)
			break
		}
	}
}

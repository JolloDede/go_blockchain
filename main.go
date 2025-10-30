package main

import (
	"fmt"

	"github.com/JolloDede/go_blockchain/pkg"
)

func main() {
	chain := pkg.CreateBlockchain()

	transaction := &pkg.Transaction{Sender: "User 1", Reciever: "User 2", Amount: 10}
	newBlock := pkg.CreateBlock(chain.Chain[len(chain.Chain)-1].Hash, []*pkg.Transaction{transaction})
	chain.AddBlock(newBlock)

	for _, block := range chain.Chain {
		fmt.Printf("Hash: %s \n", block.Hash)
	}
}

package main

import (
	"fmt"

	"github.com/JolloDede/go_blockchain/pkg"
)

func main() {
	chain := pkg.CreateBlockchain()

	chantal := pkg.CreateWallet("Chantal")
	bob := pkg.CreateWallet("Bob")

	transaction := chantal.MakeTransaction(bob.PublicKey, 10)

	newBlock := pkg.CreateBlock(chain.Chain[len(chain.Chain)-1].Hash, []*pkg.Transaction{transaction})
	chain.AddBlock(newBlock)

	for _, block := range chain.Chain {
		fmt.Printf("Hash: %s \nNonce: %d\n", block.Hash, block.Nonce)
	}
}

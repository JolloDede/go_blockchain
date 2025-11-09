package main

import (
	"crypto/rsa"
	"fmt"

	"github.com/JolloDede/go_blockchain/pkg"
)

func main() {
	chain := pkg.CreateBlockchain()
	fiendsList := []*rsa.PublicKey{}

	chantal := pkg.NewUser("Chantal", "Chantal loves to have blockchain assets", chain)
	chantal.AddWallet(pkg.CreateWallet("main"))

	bob := pkg.NewUser("Bob", "Bob loves to mine blocks", chain)
	fiendsList = append(fiendsList, bob.AddWallet(pkg.CreateWallet("Bob")))

	transaction := chantal.MakeTransaction(fiendsList[0], 10)

	bob.MineBlock([]*pkg.Transaction{transaction})

	for _, block := range chain.Chain {
		fmt.Printf("Hash: %s \nNonce: %d\n", block.Hash, block.Nonce)
	}
}

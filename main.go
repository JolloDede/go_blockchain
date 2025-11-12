package main

import (
	"fmt"

	"github.com/JolloDede/go_blockchain/pkg/blockchain"
	"github.com/JolloDede/go_blockchain/pkg/user"
)

func main() {
	chain := blockchain.CreateBlockchain()

	chantal := user.NewUser("Chantal", "Chantal loves to have blockchain assets", chain)
	chantal.AddWallet(chain.AddWallet())

	bob := user.NewUser("Bob", "Bob loves to mine blocks", chain)
	bob.AddWallet(chain.AddWallet())

	chantal.AddFriend(bob.GivePublicKey())
	bob.AddFriend(chantal.GivePublicKey())

	go chantalsLive(chantal)
	go bobsLive(bob)

	for _, block := range chain.Chain {
		fmt.Printf("Hash: %s \nNonce: %d\n", block.Hash, block.Nonce)
	}
}

func chantalsLive(u *user.User) {
	bobsIndex := 0 // index of bob in friends list
	u.MakeTransaction(u.GetFriend(bobsIndex), 10.0)
}

func bobsLive(u *user.User) {
	chantalsIndex := 0 // index of chantal in friends list
	u.MineBlock()

	u.MakeTransaction(u.GetFriend(chantalsIndex), 5.0)
	u.MineBlock()
}

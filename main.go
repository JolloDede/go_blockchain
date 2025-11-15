package main

import (
	"fmt"
	"sync"

	"github.com/JolloDede/go_blockchain/pkg/blockchain"
	"github.com/JolloDede/go_blockchain/pkg/user"
)

func main() {
	var wg sync.WaitGroup
	chain := blockchain.CreateBlockchain()

	chantal := user.NewUser("Chantal", "Chantal loves to have blockchain assets", chain)
	chantal.AddWallet(chain.AddWallet())

	bob := user.NewUser("Bob", "Bob loves to mine blocks", chain)
	bob.AddWallet(chain.AddWallet())

	chantal.AddFriend(bob.GivePublicKey())
	bob.AddFriend(chantal.GivePublicKey())

	wg.Add(2)
	go func() {
		defer wg.Done()
		chantalsLive(chantal)
	}()
	go func() {
		defer wg.Done()
		bobsLive(bob)
	}()
	wg.Wait()

	for _, block := range chain.Chain {
		fmt.Printf("Hash: %s \nNonce: %d\n", block.Hash, block.Nonce)
	}
}

func chantalsLive(u *user.User) {
	bobsIndex := 0 // index of bob in friends list
	err := u.MakeTransaction(u.GetFriend(bobsIndex), 10.0)

	if err != nil {
		println("Chantal's transaction failed:", err.Error())
	}

	u.MineBlock()
}

func bobsLive(u *user.User) {
	chantalsIndex := 0 // index of chantal in friends list
	u.MineBlock()

	err := u.MakeTransaction(u.GetFriend(chantalsIndex), 5.0)

	if err != nil {
		println("Bob's transaction failed:", err.Error())
	}

	u.MineBlock()
}

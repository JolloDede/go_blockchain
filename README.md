# Go blockchain

A simple blockchain implemenation with go.

## Table of Contents

- [Description](#description)
- [Inner Workings](#inner-workings-of-the-project)
- [How to run](#how-to-run)
- [Future development](#future-development)
- [Authors](#authors)

## Description 

This project is a exploration of blockchains.


## Inner workings of the project

### Blockchain

The main structure is, for simplicity, a slice of [Blocks](#block). Managing the chain and storing all the [transactions](#transaction) not processed by a miner. When Transaction get added to the chain the [wallet](#wallet) balance, of both recipient and sender, is beeing updated.

### Block

This type is grouping the [transactions](#transaction) and is a element of the chain. The chain is like a linked list. The blocks are linked to each other by having a hash as a identifier and the hash of the previos block.

### Transaction

Structure for holding a payment from a sender to a reciever. The sender an the reciever are beeing represented as a PublicKey. The amount is a float64.

### User

This struct is representing a person which acts on the [blockchain](#blockchain). Creating a [wallet](#wallet), Making [transactions](#transaction) or mining.

### Wallet

Wallets hold a balance and the PrivateKey. The PublicKey is accessible trough a getter function. A [user](#user) can have a wallet and send and recieve money trough it.

## How to run

A prerequisite is having [go](https://go.dev/) installed.

With go installed you can just clone the repository:
```sh
git clone https://github.com/JolloDede/go_blockchain.git
```

Change into the directory and run the project
```sh
cd go_blockchain
go run main.go
```

## Future development

- Blockchain has a map of blocks with the hash as their identifier
  - Getting the latest block would then also need some more Logik (Nonce + Latest)

- Transaction could have the amount as a currency field
  - The sent amounts may to to low for the currency

## Authors

- [@JolloDede](https://github.com/JolloDede)

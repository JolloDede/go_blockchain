package pkg

// Transaction is the type of data that we whant to store in our chain.
type Transaction struct {
	Sender   string
	Reciever string
	Amount   float64
}

// The blockchain constructor.
// Blockchain is initialized with the genesis block
func CreateBlockchain() *Blockchain {
	bc := &Blockchain{Chain: []*Block{CreateGenesisBlock()}}
	return bc
}

// The blockchain containing a slice of all the blocks
type Blockchain struct {
	Chain []*Block
}

// Function that takes in a block and calcucaltes a proof of work before adding the block to the chain
func (bc *Blockchain) AddBlock(b *Block) {
	b.PrevHash = bc.Chain[len(bc.Chain)-1].Hash

	const difficulty = 4
	b.ProofOfWork(difficulty)
	bc.Chain = append(bc.Chain, b)
}

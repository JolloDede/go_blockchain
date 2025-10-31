package pkg

type Transaction struct {
	Sender   string
	Reciever string
	Amount   float64
}

func CreateBlockchain() *Blockchain {
	bc := &Blockchain{Chain: []*Block{CreateGenesisBlock()}}
	return bc
}

type Blockchain struct {
	Chain []*Block
}

func (bc *Blockchain) AddBlock(b *Block) {
	b.PrevHash = bc.Chain[len(bc.Chain)-1].Hash

	const difficulty = 4
	b.ProofOfWork(difficulty)
	bc.Chain = append(bc.Chain, b)
}

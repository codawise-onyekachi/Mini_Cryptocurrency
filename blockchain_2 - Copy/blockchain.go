package main

type Block struct {
	// fields for the block data, a timestamp, a reference to the previous block, and a cryptographic hash of the block's contents
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(newBlock *Block) {
	if len(bc.blocks) == 0 {
		bc.blocks = append(bc.blocks, newBlock)
		return
	}

	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock.PrevBlockHash = prevBlock.Hash
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{}}
}

func main() {
	bc := NewBlockchain()
	// Add blocks to the blockchain
}

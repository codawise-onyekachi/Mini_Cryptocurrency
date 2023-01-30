package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.setHash()
	return block
}

type Blockchain struct {
	blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewBlock("Genesis Block", []byte{})}}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

const difficulty = 2

func proofOfWork(block *Block, difficulty int) {
	for i := 0; ; i++ {
		block.setHash()
		var intTimestamp int64
		intTimestamp, err := strconv.ParseInt(string(rune(block.Timestamp)), 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		if bytes.Equal(block.Hash[:difficulty], make([]byte, difficulty)) {
			fmt.Println(block.Hash)
			break
		}
		block.Timestamp = intTimestamp
	}
}

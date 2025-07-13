package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Block structure
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) GenerateHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	block_headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(block_headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.GenerateHash()
	return block
}

/*
To add a new block we need an existing block, but there’re not blocks in our blockchain!
So, in any blockchain, there must be at least one block,
and such block, the first in the chain, is called genesis block. Let’s implement a method that creates such a block:
*/

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

type Blockchain struct {
	blocks []*Block
}

func (bchain *Blockchain) AddBlock(data string) {
	prevBlock := bchain.blocks[len(bchain.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bchain.blocks = append(bchain.blocks, newBlock)
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bchain := NewBlockChain()

	bchain.AddBlock("Lit won 10k")
	bchain.AddBlock("Lit lost 5k")

	for _, block := range bchain.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}

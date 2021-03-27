package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	timestamp    string
	data         string
	previousHash string
	hash         string
}

func (b *Block) calculateHash() {
	h := sha256.New()
	h.Write([]byte(b.previousHash + b.timestamp + b.data))
	b.hash = string(h.Sum(nil))
}

func createBlock(data string, prevHash string) *Block {
	block := &Block{
		timestamp:    time.Now().Format("2006-01-02"),
		data:         data,
		previousHash: prevHash,
	}
	block.calculateHash()
	return block
}

type BlockChain struct {
	blocks []*Block
}

func createGenesisBlock() *Block {
	return createBlock("Genesis", "")
}

func (bc *BlockChain) addBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := createBlock(data, prevBlock.hash)
	bc.blocks = append(bc.blocks, block)
}

func initBlockChain() *BlockChain {
	return &BlockChain{[]*Block{createGenesisBlock()}}
}

func main() {
	chain := initBlockChain()
	chain.addBlock("test blockchain")
	fmt.Printf("Previous hash : %v\n", chain.blocks[1].previousHash)
	fmt.Printf("Hash : %v\n", chain.blocks[1].hash)
	fmt.Printf("Data : %v\n", chain.blocks[1].data)
	fmt.Printf("Date : %v\n", chain.blocks[1].timestamp)
}

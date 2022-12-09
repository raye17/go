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
	Hash          []byte
	PrevBlockHash []byte
}
type Blockchain struct {
	blocks []*Block
}

func main() {
	bc := NewBlockchain()
	bc.AddBlock("send 1 BTC to raye")
	bc.AddBlock("send 4 BTC to raye")
	for _, block := range bc.blocks {
		fmt.Printf("Prev.hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

// SetHash 计算哈希
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// NewBlock 生成新的块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		[]byte{},
		prevBlockHash,
	}
	block.SetHash()
	return block

}
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

//func isBlockValid(newBlock, oldBlock Block) bool {
//	if oldBlock.Index+1 != newBlock.Index {
//		return false
//	}
//	if oldBlock.Hash != newBlock.PrevHash {
//		return false
//	}
//	if calculateHash(newBlock) != newBlock.Hash {
//		return false
//	}
//	return true
//}
//func replaceChain(newBlocks []Block) {
//	if len(newBlocks) > len(Blockchain) {
//		Blockchain = newBlocks
//	}
//}

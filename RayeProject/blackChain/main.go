package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"time"
)

// 定义挖矿难度
const targetBits = 32

var maxNonce = math.MaxInt64

// Block 定义一个块
type Block struct {
	Timestamp     int64
	Data          []byte
	Hash          []byte
	PrevBlockHash []byte
	Nonce         int
}
type Blockchain struct {
	blocks []*Block
}
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func main() {
	bc := NewBlockchain()
	bc.AddBlock("send 1 BTC to raye")
	bc.AddBlock("send 2 BTC to raye")
	bc.AddBlock("send 3 BTC to raye")
	bc.AddBlock("send 4 BTC to raye")
	bc.AddBlock("send 5 BTC to raye")
	bc.AddBlock("send 6 BTC to raye")
	for _, block := range bc.blocks {
		fmt.Printf("Prev.hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
		pow := NewProofOfWork(block)
		fmt.Printf("PoW:%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

// NewBlock 生成新的块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		[]byte{},
		prevBlockHash,
		0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block

}

// NewGenesisBlock 生成创世块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain 创建区块链
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r成功挖矿! %x", hash)
			break
		} else {
			fmt.Printf("\r正在挖矿中...%x", hash)
			nonce++
		}
	}
	fmt.Println()
	fmt.Println(nonce)
	return nonce, hash[:]
}
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

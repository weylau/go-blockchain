package core

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	blockChain := &BlockChain{}
	newBlock := blockChain.CreateGenisisBlock()
	blockChain.addBlock(newBlock)
	return blockChain
}

func (bc *BlockChain) SendData(data string) (err error) {
	newBlock := &Block{}
	if bc.GetBlockChainSize() <= 0 {
		return errors.New("blockchain had not genisis block")
	}
	newBlock = bc.CreateNewBlock(bc.GetLastBlock(), data)
	if !bc.IsVaild(newBlock) {
		return errors.New("block is vaild fail")
	}
	bc.addBlock(newBlock)
	return nil
}

func (bc *BlockChain) addBlock(block *Block) {
	if bc.GetBlockChainSize() == 0 {
		bc.Blocks = append(bc.Blocks, block)
		return
	}
	if bc.IsVaild(block) {
		bc.Blocks = append(bc.Blocks, block)
	}
}

func (bc *BlockChain) IsVaild(block *Block) bool {
	lastBlock := bc.GetLastBlock()
	if (block.Index - 1) != lastBlock.Index {
		fmt.Println("block.Index != lastBlock.Index")
		return false
	}
	if block.PrevBlockHash != lastBlock.Hash {
		fmt.Println("block.PrevBlockHash != lastBlock.Hash")
		return false
	}
	newBlockHash := MakeBlockHash(block)

	if block.Hash != newBlockHash {
		fmt.Printf("block.Hash(%s) != newBlockHash(%s)", block.Hash, newBlockHash)
		return false
	}
	return true
}

func (bc *BlockChain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *BlockChain) GetBlockChainSize() int {
	return len(bc.Blocks)
}

func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("=========Block %d ========\n", block.Index)
		fmt.Printf("Timestamp：%d\n", block.Timestamp)
		fmt.Printf("PrevBlockHash：%s\n", block.PrevBlockHash)
		fmt.Printf("Hash：%s\n", block.Hash)
		fmt.Printf("Data：%s\n", block.Data)
		fmt.Println()
	}
}

func (bc *BlockChain) CreateNewBlock(prevBlock *Block, data string) (block *Block) {
	block = &Block{
		Index:         prevBlock.Index + 1,
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlock.Hash,
		Data:          data,
	}
	hash := MakeBlockHash(block)
	block.Hash = hash
	return block
}

//生成创世区块
func (bc *BlockChain) CreateGenisisBlock() (block *Block) {
	preBlock := &Block{
		Index: -1,
		Hash:  "",
	}
	block = bc.CreateNewBlock(preBlock, "fist block")
	return block
}

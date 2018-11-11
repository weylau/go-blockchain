package core

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Index         int64  `json:"index"`           //区块编号
	Timestamp     int64  `json:"timestamp"`       //区块时间
	PrevBlockHash string `json:"prev_block_hash"` //上一个区块的hash
	Hash          string `json:"hash"`            //区块hash值
	Data          string `json:"data"`            //区块数据
}

func MakeBlockHash(block *Block) (blockHash string) {
	blockData := MakeBlockData(block)

	hashByte := sha256.Sum256(blockData)
	return hex.EncodeToString(hashByte[:])
}

func MakeBlockData(block *Block) (data []byte) {
	blockData := block.Data + block.PrevBlockHash + string(block.Timestamp) + string(block.Index)
	return []byte(blockData)
}

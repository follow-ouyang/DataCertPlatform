package blockchain

import "time"

//定义区块结构体，用于表示区块
type Block struct {
	Height int64 //表示区块的高度，就是第几个区块
	TimeStamp int64 //区块产生的时间戳
	PrevHash []byte //前一个区块的hash
	Data []byte //数据字段
	Hash []byte //当前区块的hash值
	Version string //版本号
}

/**
 *创建一个新区块
 */
func NewBlock(height int64,prevHash,data []byte ) Block {
	bolck := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
		Version:   "0x01",
	}
	//block.Hash =
	return bolck
}

/**
 * 创建创世区块
 */
func CreateGenesisBlock() Block {
	 genesisblock := NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
	 return genesisblock
}
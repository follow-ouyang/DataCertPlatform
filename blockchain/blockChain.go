package blockchain

import "github.com/bolt"

const BLOCKCHAIN = "blockchain.db"
const BUCKET_NAME  = "blocks"
const LAST_HASH  = "lasthash"
/*
区块链结构体的定义，代表的是整条区块链
 */

type BlockChain struct {
	LastHash []byte//表示区块链中最新区快的哈希，用于查找最新的区块内容
	BoltDb *bolt.DB//区块链中操作区块数据文件的数据库操作对象

}

/*
创建一个区块链
 */
func NewBlockChain() BlockChain {
	var bc BlockChain
	//1、先打开文件
	db,_ := bolt.Open(BLOCKCHAIN,0600,nil)
	//2、查看chain.db文件
	db.Update(func(tx *bolt.Tx) error {
		//假设有桶
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {//说明没有桶，要创建新桶
			bucket,err := tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		lastHash := bucket.Get([]byte(LAST_HASH))
		if len(lastHash) == 0 { //说明桶当中没有lasthash记录，需要新建创世区块
			//创世区块
			genesis := CreateGenesisBlock()
			//区块序列化以后的数据
			gensisBytes := genesis.Serialize()
			//创世区块保存到boltdb中
			bucket.Put(genesis.Hash, gensisBytes)
			//更新指向最新区块的lasthash的值
			bucket.Put([]byte(LAST_HASH), genesis, nil)
			bc := BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
		}else {//桶中已有lasthash的记录，不再需要创世区块，只需要读取即可
			lasthash := bucket.Get([]byte(LAST_HASH))
			bc = BlockChain{
				LastHash: lasthash,
				BoltDb:   db,
			}
		}
		return nil
	})
	return bc

	//db,err := bolt.Open(BLOCKCHAIN,0600,nil)
	//if err != nil {
	//	panic(err.Error())
	//}
	//bc := BlockChain{
	//	LastHash: genesis.Hash,
	//	BoltDb:   db,
	//}
	//把创世区块保存到数据库文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket,err := tx.CreateBucket([]byte(BUCKET_NAME))
		if err != nil {
			panic(err.Error())
		}
		//序列化
		genesisBytes := genesis.Serialize()
		//把创世区块存储到桶中
		bucket.Put(genesis.Hash,genesisBytes)
		//更新最新区快的哈希值
		bucket.Put([]byte(LAST_HASH),genesis.Hash)
		return nil
	})
	return bc
}

/*
保存数据到区块链中：先生成一个区块添加到区块链中
 */
func (bc BlockChain) SaveBlock(data []byte) {
	//1、从文件中读取到最新的区块
	db := bc.BoltDb
	var lastBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("读取区块数据失败")
		}
		//lastHash := bucket.Get([]byte(LAST_HASH))
		lastBlockBytes := bucket.Get(bc.LastHash)
		//反序列化
		lastBlock,_ = DeSerialize(lastBlockBytes)
		return nil
	})

	//新建一个区块
	newBlock := NewBlock(lastBlock.Height+1,lastBlock.Hash,data)
	//把新区块存到文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//把新创建的区块存入到boltdb数据库中
		bucket.Put(newBlock.Hash,newBlock.Serialize())
		//更新LASTHASH对应的值，更新为最新存储的区块的hash值
		bucket.Put([]byte(LAST_HASH),newBlock.Hash)
		//将区块链实例的LASHASH值更新为最新区快的哈希
		bc.LastHash = newBlock.Hash
		return nil

	})

}
package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	block0 := blockchain.CreateGenesisBlock()//创建创世区块
	block1 := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte{})
	//block2 := blockchain.NewBlock(block1.Height+1,block1.Hash,[]byte{})
	fmt.Printf("block0的哈希：%x\n",block0.Hash)
	fmt.Printf("block1的哈希：%x\n",block1.Hash)
	//fmt.Printf("block2的哈希：%x\n",block2.Hash)
	fmt.Printf("block1的PrevHash：%x\n",block1.PrevHash)

	block0Bytes := block0.Serialize()
	fmt.Println("创世区块通过gob序列化后：",block0Bytes)
	deblock0,err := blockchain.DeSerialize(block0Bytes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("反序列化后区块的高度是：",deblock0.Height)
	fmt.Printf("反序列化后的区块的哈希：%x\n",deblock0.Hash)

	return
	//序列化:将数据从内存中形式转换为可以持久化存储在硬盘上或者网络上传输的形式，称之为序列化
	//反序列化：将数据从文件中或者网络中读取，然后转化到计算机内存中的过程
	//序列化和反序列化有很多种方式
	//        json、xml:
	//				序列化：Marshal
	//				反序列化：UnMarshal
	//只有序列化以后的文件才能进行传输
	blockJson,_ := json.Marshal(block0)
	fmt.Println("通过json序列化以后的block：",string(blockJson))

	//连接数据库
	db_mysql.Connect()
	//设置静态资源文件映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}


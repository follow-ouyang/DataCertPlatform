package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//1、创世区块
	bc := blockchain.NewBlockChain()//封装
	fmt.Printf("创世区块的哈希值：%x\n",bc.LastHash)
	bc.SaveBlock([]byte("用户要保存到区块中的数据"))

	return
	//序列化:将数据从内存中形式转换为可以持久化存储在硬盘上或者网络上传输的形式，称之为序列化
	//反序列化：将数据从文件中或者网络中读取，然后转化到计算机内存中的过程
	//序列化和反序列化有很多种方式
	//        json、xml:
	//				序列化：Marshal
	//				反序列化：UnMarshal
	//只有序列化以后的文件才能进行传输

	//连接数据库
	db_mysql.Connect()
	//设置静态资源文件映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}


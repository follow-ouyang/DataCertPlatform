package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"github.com/astaxie/beego"
)

func main() {
	/*
	测试代码
	 */
	////1、创世区块
	//bc := blockchain.NewBlockChain()//封装
	//bc.SaveData([]byte("别惹我，我来大姨夫了"))
	//blocks,err := bc.QueryAllBlocks()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//for _, block := range blocks {
	//	fmt.Printf("高度：%d，区块内数据：%s\n",block.Height,block.Data)
	//}
	//
	//return
	//序列化:将数据从内存中形式转换为可以持久化存储在硬盘上或者网络上传输的形式，称之为序列化
	//反序列化：将数据从文件中或者网络中读取，然后转化到计算机内存中的过程
	//序列化和反序列化有很多种方式
	//        json、xml:
	//				序列化：Marshal
	//				反序列化：UnMarshal
	//只有序列化以后的文件才能进行传输

	//先准备一条区块链
	blockchain.NewBlockChain()

	//连接数据库
	db_mysql.Connect()
	//设置静态资源文件映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}


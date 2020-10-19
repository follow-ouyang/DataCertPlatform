package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	block0 := blockchain.CreateGenesisBlock()
	block1 := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte("a"))
	//block2 := blockchain.NewBlock(block1.Height+1,block1.Hash,[]byte())
	fmt.Println(block0)
	fmt.Println(block1)
	//连接数据库
	db_mysql.Connect()
	//设置静态资源文件映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}


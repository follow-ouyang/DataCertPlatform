package controllers

import (
	"DataCertPlatform/blockchain"
	"fmt"
	"github.com/astaxie/beego"
)

type CertDataController struct {
	beego.Controller
}

/*
该get方法用于处理浏览器的get请求，往查看证书详情页面跳转
 */
func (c *CertDataController) Get() {
	//1、接受和解析前端页面传递的数据
	cert_id := c.GetString("cert_id")

	//2、到区块链上查询区块数据
	block,err := blockchain.CHAIN.QUeryBlockByCertId(cert_id)
	if err != nil {
		c.Ctx.WriteString("抱歉，查询脸上数据遇到错误，请重试")
		return
	}
	if block == nil { //遍历整条区块链，但未查到数据
		c.Ctx.WriteString("抱歉，脸上没有你的东西哦")
		return
	}
	fmt.Println("查询到的区块的高度为：",block.Height)

	c.Data["CertId"] = string(block.Data)
	//3、跳转证书详情页面
	c.TplName = "cert_detail.html"
}

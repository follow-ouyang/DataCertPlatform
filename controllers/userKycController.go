package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
)

type UserKycController struct {
	beego.Controller
}

func (u *UserKycController) Get() {
	u.TplName = "user_kyc.html"
}

/*
form表单的post请求，用于处理实名认证业务
 */
func (u *UserKycController) Post() {
	//1、解析前端的数据
	var user models.User
	err := u.ParseForm(&user)
	if err != nil {
		u.Ctx.WriteString("抱歉，解析错误，请重试")
		return
	}
	if user.Name == "" || user.Card == "" || user.Sex == "" {
		u.Ctx.WriteString("对不起，您的信息未输入完全，请重新输入")
		return
	}
	//2、把用户的实名认证更新到数据库的用户表中
	_,err = user.UpdateUser()
	//3、判断实名认证的结果
	if err != nil {
		u.Ctx.WriteString("抱歉，实名认证失败，请重试")
		return
	}

	//4、跳转到上传页面
	u.TplName = "home.html"

}
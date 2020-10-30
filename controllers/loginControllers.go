package controllers

import (
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type LoginController struct {
	beego.Controller
}


//Post方法处理用户登录请求
func (l *LoginController) Post() {
	//1、解析客户端用户提交的登陆数据
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("抱歉，用户登录信息解析失败，请重试")
		fmt.Println(err.Error())
		return
	}

	//对用户输入的信息进行判断
	if len(user.Phone) != 11 {
		l.Ctx.WriteString("手机号码格式输入错误，请重新输入并登录。。。")
		return
	}
	if len(user.Password) >18 || len(user.Password) < 6 {
		l.Ctx.WriteString("密码格式输入错误，请重新输入并登录。。。")
		return
	}

	//2、根据解析到的数据，执行数据库查询操作
	u,err := user.QueryUser()

	//3、判断数据库的查询结果
	if err != nil {
		l.Ctx.WriteString("抱歉，用户登录失败，请重试")
		fmt.Println(err.Error())
		return
	}

	//3、1  增加逻辑，判断用户是否已经实名认证，如果没有则跳转到实名认证页面
	if strings.TrimSpace(u.Name) == "" || strings.TrimSpace(u.Card) == "" {
		l.Data["Name"] = u.Name
		l.TplName = "user_kyc.html"
		return
	}


	//4、根据查询结果返回客户端相应的信息或页面跳转
	l.Data["Phone"] = u.Phone//动态数据设置
	l.TplName = "home.html"
}

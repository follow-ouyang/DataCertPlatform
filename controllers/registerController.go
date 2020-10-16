package controllers

import (
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

/**
该方法用于处理用户注册的逻辑
 */
func (r *RegisterController) Get() {
	r.TplName = "login.html"

}

func (r *RegisterController) Post() {
	//1 解析用户端提交的请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉数据解析失败，请重新尝试。。。")
		return
	}
	//对提交的数据进行格式判断
	if len(user.Phone) != 11 {
		r.Ctx.WriteString("手机号码格式输入错误，请重新输入并注册。。。")
		return
	}
	if len(user.Password) >18 || len(user.Password) < 6 {
		r.Ctx.WriteString("密码格式输入错误，请重新输入并注册。。。")
		return
	}

	//2 将解析到的数据保存到数据库中
	row,err := models.User.AddUser(user)
	if err != nil {
		r.Ctx.WriteString("数据保存失败，请重试。。。。")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(row)

	//3 将处理结果返回给客户端浏览器
	//3.1如果成功，跳转登录页面
	//tpl：template 模板
	r.TplName = "login.html"
	//3.2如果失败，提示错误信息

}
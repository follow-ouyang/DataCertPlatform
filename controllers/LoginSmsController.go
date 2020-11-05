package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
)

type LoginSmsController struct {
	beego.Controller
}

/*
浏览器发起的短信验证码的登录请求
 */
func (l *LoginSmsController) Get() {
	l.TplName = "login_sms.html"
}
//短信验证登录功能
func (l *LoginSmsController) Post() {
	//biz_id,phone,code
	var smsLogin models.SmsLogin
	err := l.ParseForm(&smsLogin)
	if err != nil {
		l.Ctx.WriteString("抱歉，验证码登陆失败，请重试")
	}

	//那用户的

}
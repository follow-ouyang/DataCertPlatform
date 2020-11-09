package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
	"time"
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
		return
	}
	//1、先拿手机号，查询user表，看用户是否已注册
	user := models.User{
		Phone:smsLogin.Phone,
	}
	u,err := user.QueryUserByPhone()
	if err != nil {
		l.Ctx.WriteString("对八起，验证码登陆失败，请重试")
		return
	}
	if u.Phone == "" { //说明在用户表中不存在该记录，即未注册
		l.Ctx.WriteString("抱歉,您还未注册，请先注册")
		return
	}

	//拿用户提交的登录信息到数据库中查询
	sms,err := models.QuerySmsRecord(smsLogin.BizId,smsLogin.Phone,smsLogin.Code)
	if err != nil {
		l.Ctx.WriteString("抱歉，验证码登录遇到错误，请重试")
		return
	}
	if sms.BizId == "" {//验证码错误，手机号错误
		l.Ctx.WriteString("手机号或验证码错误，请重新输入")
		return
	}
	now := time.Now().Unix()
	if (now - sms.TimeStamp) > 30000 {
		l.Ctx.WriteString("验证码已失效，请重新获取")
		return
	}
	//验证码正常，跳转主界面
	l.Data["Phone"] = smsLogin.Phone
	l.TplName = "home.html"

}
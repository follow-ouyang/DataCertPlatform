package controllers

import (
	"DataCertPlatform/models"
	"DataCertPlatform/utils"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type SendSmsController struct {
	beego.Controller
}

/*
发送验证码的功能
 */
func (s *SendSmsController) Post() {
	fmt.Println("发送成功")
	var smsLogin models.SmsLogin
	err := s.ParseForm(&smsLogin)
	if err != nil {
		s.Ctx.WriteString("发送验证码数据解析失败")
		return
	}
	phone := smsLogin.Phone
	code := utils.GenRandCode(6)//返回一个6位数的随机数
	result,err := utils.SendSms(phone,code,utils.SMS_TLP_REGISTER)
	if err != nil {
		s.Ctx.WriteString("发送验证码失败，请重试")
	}
	if len(result.BizId) == 0 {
		s.Ctx.WriteString(result.Message)
		return
	}
	//到此，验证码发送成功
	smsRecord := models.SmsRecord{
		BizId:     result.BizId,
		Phone:     phone,
		Code:      code,
		Status:    result.Code,
		Message:   result.Message,
		TimeStamp: time.Now().Unix(),
	}
	_,err =  smsRecord.SaveSmsRecord()
	if err !=nil {
		s.Ctx.WriteString("抱歉，获取验证码失败，请重试")
		return
	}
	//保存成功   bizId返回前端
	s.Data["Phone"] = smsLogin.Phone
	s.Data["BizId"] = smsRecord.BizId
	//验证码登录
	s.TplName = "login_sms.html"

}

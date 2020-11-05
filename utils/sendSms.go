package utils

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego"
	"math/rand"
	"strings"
	"time"
)

type SmsCode struct {
	Code string`json:"code"`
}
type SmsResult struct {
	BizId     string
	Code      string
	Message   string
	RequestId string
}

const SMS_TLP_REGISTER = "SMS_205393604"//注册业务的短信模板
const SMS_TLP_LOGIN  = "SMS_205398654"//用户登陆的短信模板
const SMS_TLP_KYC  = ""//实名认证的短信模板

/*
该函数用于发送一条短信息
参数：
	phone:电话，接受验证码的号码
	code：发送的验证码数字
	template：模板类型
 */
func SendSms(phone string,code string,template string) (*SmsResult,error) {
	config := beego.AppConfig
	//获取配置文件中的数据
	smsSccessKey := config.String("sms_sccess_key")
	smsSccessSecret := config.String("sms_acces_secret")
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", smsSccessKey, smsSccessSecret)
	if err != nil {
		return nil,err
	}
	//batch：批量。批次
	request := dysmsapi.CreateSendSmsRequest()
	request.PhoneNumbers = phone//指定要发送的目标手机号
	request.SignName = "真的来啦"
	request.TemplateCode = template//指定短信模板
	smsCode := SmsCode{
		Code:code,
	}
	smsbytes,_ := json.Marshal(smsCode)
	request.TemplateParam = string(smsbytes)

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	//Biz : business,商业，业务
	smsResult := &SmsResult{
		BizId:     response.BizId,
		Code:      response.Code,
		Message:   response.Message,
		RequestId: response.RequestId,
	}
	return smsResult,nil

}

func GenRandCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0;i<width ; i++ {
		fmt.Fprintf(&sb,"%d",numeric[rand.Intn(r)])
	}
	return sb.String()

}
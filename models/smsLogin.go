package models

type SmsLogin struct {
	BizId string`form:"bizId"`
	Phone string`form:"phone"`
	Code string`form:"code"`
}
package models

import "DataCertPlatform/db_mysql"

type SmsRecord struct {
	BizId     string
	Phone     string
	Code      string
	Status    string
	Message   string
	TimeStamp int64
}

/*
根据用户提交的手机号和短信验证码查询验证码是否正确及正常
 */
func QuerySmsRecord(bizId string,phone string,code string)  {

}

/*
向数据库当中插入验证码记录，该记录有阿里云第三方返回
*/
func (s SmsRecord) SaveSmsRecord() (int64, error) {
	rs,err := db_mysql.Db.Exec("insert into sms_record(biz_id,pgone,code,status,message,timestamp )value (?,?,?,?,?,?)",
		s.BizId, s.Phone,s.Code,s.Status,s.Message,s.TimeStamp)
	if err != nil {
		return -1,err
	}
	return rs.RowsAffected()
}
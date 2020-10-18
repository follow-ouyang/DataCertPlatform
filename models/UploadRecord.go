package models

import (
	"DataCertPlatform/db_mysql"
	"fmt"
)

/*
上传文件的记录
 */
type UploadRecord struct {
	Id int
	UserId int
	FileName string
	FileSize int64
	FileCert string
	FileTitle string
	CertTime int64
}

/*
把一条认证数据保存到数据库表中
 */
func (u UploadRecord) SaveRedcord() (int64,error) {
	result,err := db_mysql.Db.Exec("insert into upload_record (user_id,file_name,file_size,file_cert,file_title,cert_time)" +
		"values(?,?,?,?,?,?)",u.UserId,u.FileName,u.FileSize,u.FileCert,u.FileTitle,u.CertTime)
	if err != nil {
		fmt.Println("是这里吗：",err.Error())
		return -1,err
	}

	id,err := result.RowsAffected()
	if err != nil {
		fmt.Println("还是这里：",err.Error())
		return -1,err
	}

	return id,err

}

/*
根据用户Id查询符合条件的认证数据记录
 */
func QueryRecordsUserId(userId int) ([]UploadRecord,error) {
	rs,err :=  db_mysql.Db.Query("select id, user_id, file_name, file_size, file_cert, file_title, cert_time from upload_record where user_id = ?",userId)
	if err != nil {
		return nil,err
	}
	//从rs中读取查询到的数据，并返回
	records := make([]UploadRecord,0)//容器
	for rs.Next() {
		var record UploadRecord
		err := rs.Scan(&record.Id,&record.UserId,&record.FileName,&record.FileSize,&record.FileCert,&record.FileTitle,&record.CertTime)
		if err != nil {
			return nil,err
		}
		records = append(records,record)
	}
	return records,nil
}
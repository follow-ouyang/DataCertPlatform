package db_mysql

import (
	"DataCertPlatform/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "go-sql-driver/mysql"
)

var Db *sql.DB

//连接数据库
func Connect()  {
	config := beego.AppConfig

	////获取配置选项
	dbDriver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbpassword := config.String("db_password")
	dbIP := config.String("db_ip")
	dbName := config.String("db_name")
	//fmt.Println(dbDriver,dbUser,dbpassword)

	connUrl := dbUser+":"+dbpassword+"@tcp("+dbIP+")/"+dbName+"?charset=utf8"
	db,err := sql.Open(dbDriver,connUrl)
	if err != nil {//err不为nil，表示连接数据库时出现错误，程序在此中断就行，不用再执行
		panic("数据库连接失败，请重试")
	}
	Db = db
	//fmt.Println(db)
}

//将用户信息保存到数据库中
func AddUser(u models.User) (int64,error) {
	//将得到的密码进行Hash计算，得到密码Has值
	md5hash := md5.New()
	md5hash.Write([]byte(u.Password))
	passwordByte := md5hash.Sum(nil)
	u.Password = hex.EncodeToString(passwordByte)

	result,err := Db.Exec("insert into approve(phone,password)" +
		"value(?,?)",u.Phone,u.Password)
	if err != nil {
		fmt.Println(err.Error())
		return -1,err
	}
	row,err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return -1,err
	}
	fmt.Println(row)
	return row,err


}
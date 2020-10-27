package db_mysql

import (
	"database/sql"
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




package models

import (
	"DataCertPlatform/db_mysql"
	"DataCertPlatform/utils"
	"fmt"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

//将用户信息保存到数据库中
func (u User) AddUser() (int64, error) {
	//将得到的密码进行Hash计算，得到密码Has值
	u.Password = utils.Md5HashString(u.Password)

	result, err := db_mysql.Db.Exec("insert into approve(phone,password)"+
		"values (?,?)", u.Phone, u.Password)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	//fmt.Println(row)

	return row, err
}

/*
 *查询用户信息
 */
func (u User) QueryUser() (*User, error) {
	u.Password = utils.Md5HashString(u.Password)

	row := db_mysql.Db.QueryRow("select phone from approve where phone = ? and password = ?",
		u.Phone, u.Password)

	err := row.Scan(&u.Phone)
	if err != nil {
		return nil, err
	}
	return &u, err
}

func (u User) QueryUserByPhone() (*User, error) {
	row := db_mysql.Db.QueryRow("select id from approve where phone = ?", u.Phone)

	err := row.Scan(&u.Id)
	if err != nil {
		return nil, err
	}
	return &u, err

}

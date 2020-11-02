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
	Name     string `form:"name"` //名字
	Card     string `form:"card"` //身份证号
	Sex      string `form:"sex"`  //性别
}

/*
该方法用于更新数据库中用户记录的实名认证信息
*/
func (u User) UpdateUser() (int64, error) {
	rs, err := db_mysql.Db.Exec("update approve set  name = ?,card = ?,sex = ? where phone = ?", u.Name, u.Card, u.Sex, u.Phone)
	if err != nil {
		return -1, err
	}
	id, err := rs.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, err

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

	row := db_mysql.Db.QueryRow("select phone, name ,card from approve where phone = ? and password = ?",
		u.Phone, u.Password)

	err := row.Scan(&u.Phone, &u.Name, &u.Card)
	if err != nil {
		return nil, err
	}
	return &u, err
}

func (u User) QueryUserByPhone() (*User, error) {
	row := db_mysql.Db.QueryRow("select id from approve where phone = ?", u.Phone)
	var user User

	err := row.Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return &user, err

}

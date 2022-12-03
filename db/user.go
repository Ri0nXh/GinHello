package db

import (
	"GinHello/model"
	"log"
)

func SaveUser(user *model.UserRegister) int64 {
	// 感觉这里设计的有点问题，应该需要返回error
	result, err := DBConn.Exec("insert into go_stu.ginhello_user(email, username, password) "+
		"values (?, ?, ?);", user.Email, user.Username, user.Password)
	if err != nil {
		log.Println("insert table ginhello_user failed, Error : ", err.Error())
	}
	userId, err := result.LastInsertId()
	if err != nil {
		log.Println("get userid error : ", err.Error())
	}
	return userId
}

func UserLogin(u *model.UserLogin) (user *model.UserInfo, err error) {
	username := u.Username
	userinfo := model.UserInfo{}
	if username != "" {
		sql := "select id,username,email from ginhello_user where username = ? and password = ?"
		err = DBConn.QueryRow(sql, username, u.Password).Scan(&userinfo.Id, &userinfo.Username, &userinfo.Email)
	} else {
		sql := "select id,username,email from ginhello_user where email = ? and password = ?"
		err = DBConn.QueryRow(sql, u.Email, u.Password).Scan(&userinfo.Id, &userinfo.Username, &userinfo.Email)
	}
	return &userinfo, err
}

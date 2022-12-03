package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var DBConn *sql.DB

func Init() (err error) {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db"))
	fmt.Println(dbUrl)
	// 注意： 不能使用 := ， 否则全局变量的DBConn 是一个空地址，而Init中的DBConn 是一个内部变量
	// 会导致连接mysql是正常的，当时执行mysql时会出错，因为全局的DBConn时一个空地址
	DBConn, err = sql.Open("mysql", dbUrl)
	if err != nil {
		return err
	}
	DBConn.SetMaxOpenConns(30)
	DBConn.SetMaxIdleConns(5)
	// 尝试连接是否正常
	err = DBConn.Ping()
	if err != nil {
		return err
	}
	return nil
}

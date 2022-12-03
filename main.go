package main

import (
	"GinHello/config"
	"GinHello/db"
	"GinHello/router"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// 初始化配置文件
	if err := config.Init(); err != nil {
		fmt.Printf("yaml Init Failed!, Error : ", err.Error())
		return
	}

	// 初始化数据库
	if err := db.Init(); err != nil {
		fmt.Printf("Mysql Init Failed!, Error : ", err.Error())
		return
	}

	router := router.SetupRouter()
	router.Run(fmt.Sprintf(":%d", viper.GetInt("server_port")))
}

package config

import (
	"github.com/spf13/viper"
)

var Conf = new(AppConf)

type AppConf struct {
	Mode         string `yaml:"mode"`
	ServerPort   int    `yaml:"server_port"`
	*MySQLConfig `yaml:"mysql"`
	*RedisConfig `yaml:"redis"`
}

type MySQLConfig struct {
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DB           string `yaml:"db"`
	Port         int    `yaml:"port"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	DB   string `yaml:"db"`
	Port int    `yaml:"port"`
}

func Init() error {

	//viper.SetConfigType("yaml")
	//viper.MergeConfig(r)
	//r, err := os.Open("config.yaml")
	// 注意此处的文件路径时以项目为根目录进行读取，即 main.go
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		return err
	}
	return err
}

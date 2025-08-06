package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Mysql struct {
		User     string
		Password string
		Host     string
		Port     int
		Database string
	}
	Redis struct {
		Host     string
		Password string
		Db       int
	}
}

var Appconf Config

func Inits() {
	// 设置配置文件名和类型
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	// 添加多个可能的配置文件路径
	viper.AddConfigPath("./basic/config/")       // 从user目录运行
	viper.AddConfigPath("./user/basic/config/")  // 从server目录运行
	viper.AddConfigPath("../user/basic/config/") // 从其他子目录运行
	viper.AddConfigPath(".")                     // 当前目录

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("配置文件读取失败: %v\n请确保在正确的目录运行程序，或检查配置文件路径", err))
	}

	fmt.Printf("配置文件读取成功，使用文件: %s\n", viper.ConfigFileUsed())

	err = viper.Unmarshal(&Appconf)
	if err != nil {
		panic(fmt.Sprintf("配置文件解析失败: %v", err))
	}

	fmt.Printf("配置信息加载完成: MySQL=%s:%d/%s, Redis=%s\n",
		Appconf.Mysql.Host, Appconf.Mysql.Port, Appconf.Mysql.Database, Appconf.Redis.Host)
}

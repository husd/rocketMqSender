package sender

import (
	"fmt"
	"github.com/spf13/viper"
)

/**
 * 读取配置文件
 * @author hushengdong
 */

var v *viper.Viper

func initConfig(f string) {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("解析配置文件错误:", err)
		return
	}
	v = viper.GetViper()

	fmt.Println("------------------基本配置信息------------------")
	fmt.Println("nameserver:", getNameserver())
	fmt.Println("broker:", getBroker())
	fmt.Println("version:", getVersion())
	fmt.Println("------------------基本配置信息------------------")
	fmt.Println("                                              ")
}

func getNameserver() string {

	return v.GetString("nameserver")
}

func getBroker() string {

	return v.GetString("broker")
}

func getVersion() int32 {

	code := v.GetInt("version")
	if code <= 0 {
		code = 79
	}
	return int32(code)
}

func getExtFieldsByCode(code reqCode) map[string]string {

	key := fmt.Sprintf("extField_%d", int(code))
	return v.GetStringMapString(key)
}

func getBodyByCode(code reqCode) string {

	key := fmt.Sprintf("body_%d", int(code))
	return v.GetString(key)
}

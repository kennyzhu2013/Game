/*
@Time : 2019/6/28 15:10
@Author : kenny zhu
@File : conf.go
@Software: GoLand
@Others:
*/
package Conf

import "github.com/spf13/viper"

func InitConfig(filename string)  {
	viper.SetConfigType("toml")
	viper.SetConfigFile(filename)
	viper.ReadInConfig()


}

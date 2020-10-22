// Time : 2020/9/19 15:04
// Author : Kieran

// config 配置，读取配置文件
// starter和一些业务会读配置，最好让starter统一去读
package config

// config.go something

import (
	"github.com/spf13/viper"
)

var C = defaultConfig()

// envConfig.go
type envConfig struct {
	Dispatch DispatchConfig
	Conn     ConnConfig
	Logic    LogicConfig
}

type DispatchConfig struct {
}

type LogicConfig struct {
	RequestTimeOut   int

	AutoReConnection bool
	ReConnectionIntervalSecond int
	ReConnectionRetryCount int
}

type ConnConfig struct {
	ReadTimeout int
}

func init() {
	viper.SetDefault("DISPATCH_CONFIG", "{}")
}

// 加载远程配置
func remoteConfig() envConfig {
	return envConfig{}
}

// setupConfig 加载本地配置
func setupConfig() envConfig {
	return envConfig{}
}

func defaultConfig() envConfig {
	return envConfig{
		Conn: ConnConfig{
			ReadTimeout: 10,
		},
		Logic: LogicConfig{
			RequestTimeOut:             10,
			AutoReConnection:           true,
			ReConnectionIntervalSecond: 2,
			ReConnectionRetryCount:     5,
		}}
}

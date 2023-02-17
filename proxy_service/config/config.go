package config

import "github.com/spf13/viper"

func InitConfig(path string, cfgName string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(cfgName)
	return viper.ReadInConfig()
}

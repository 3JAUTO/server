package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

var envConf EnvConfig

type EnvConfig struct {
	Addr     string `mapstructure:"addr"`
	LogFile  string `mapstructure:"logfile"`
	LogLevel string `mapstructure:"loglevel"`
	DBUrl    string `mapstructure:"db_url"`
	DBUser   string `mapstructure:"db_user"`
	DBPasswd string `mapstructure:"db_password"`
	DBName   string `mapstructure:"db_name"`
}

func initConfigFile() error {
	if Env == "" {
		return fmt.Errorf("server environment mustn't be empty. %s", Env)
	}
	if ConfPath == "" {
		viper.AddConfigPath(defaultConfPath)
	} else {
		viper.AddConfigPath(ConfPath)
	}
	viper.SetConfigName("conf_" + Env)
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("fail to read config file. %v", err)
	}
	if err := viper.Unmarshal(&envConf); err != nil {
		return fmt.Errorf("fail to unmarshal config file. %v", err)
	}
	return nil
}

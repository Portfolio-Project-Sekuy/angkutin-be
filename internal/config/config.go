package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

// GetConfig get config
func GetConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.AddConfigPath("./../../..")
	viper.SetConfigName("config")
	viper.SetEnvPrefix("svc")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logrus.Warningf("%v", err)
	}
}

// Env env
func Env() string {
	return viper.GetString("env")
}

// LogLevel log level
func LogLevel() string {
	return viper.GetString("log_level")
}

// HTTPPort http server port
func HTTPPort() string {
	return viper.GetString("http_port")
}

package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

// GetConfig get config
func GetConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logrus.Warningf("%v", err)
	}
}

// Env env
func Env() string {
	return viper.GetString("ENV")
}

// LogLevel log level
func LogLevel() string {
	return viper.GetString("LOG_LEVEL")
}

// Port http server port
func Port() string {
	return viper.GetString("PORT")
}

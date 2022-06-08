package config

import (
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func viperConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("Fatal error config file: %s \n", err)
		panic(err)
	}
}
package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("smtp")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "8080")

	err := viper.BindEnv("email", "email")
	checkCriticalErr(err)
	err = viper.BindEnv("pass", "pass")
	checkCriticalErr(err)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config found, using default values")
		} else {
			fmt.Println("Failed to read config file, falling back to defaults", err)
		}
		return
	}

	if viper.GetString("email") == "" {
		panic("No email supplied for smtp")
	}
	if viper.GetString("pass") == "" {
		panic("No password supplied for smtp")
	}
	fmt.Println("Config successfully loaded")
}

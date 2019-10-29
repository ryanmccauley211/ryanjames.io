package main

import (
	//"fmt"
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("smtp")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetDefault("email", "ryanmccauley211@gmail.com")
	viper.SetDefault("host", "")
	viper.SetDefault("port", "8080")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config found, using default values")
		} else {
			fmt.Printf("Failed to read config file, falling back to defaults\n", err)
		}
		return
	}
	fmt.Println("Config successfully loaded")
}

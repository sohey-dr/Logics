package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
	fmt.Println("DB_HOST:", viper.GetString("DB_HOST"))
	fmt.Println("DB_PORT:", viper.GetInt("DB_PORT"))
	fmt.Println("DB_USER:", viper.GetString("DB_USER"))
	fmt.Println("DB_PASS:", viper.GetString("DB_PASS"))
	fmt.Println("DB_NAME:", viper.GetString("DB_NAME"))
}
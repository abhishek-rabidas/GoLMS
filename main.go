package main

import (
	"github.com/spf13/viper"
	"gomvc/controller"
	"gomvc/resources"
)

func init() {
	resources.PopulateUsers()
	viper.SetConfigFile("resources/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	server := controller.NewEchoServer()
	defer server.Close()
	err := server.Start(":" + viper.GetString("server.port"))
	if err != nil {
		panic(err)
	}
}

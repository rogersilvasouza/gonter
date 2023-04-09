/*
Copyright © 2023 Roger Souza <rogersilvasouza@hotmail.com>
*/
package main

import (
	"fmt"

	"github.com/rogersilvasouza/gonter/internal/google"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("urls", []string{})
	viper.SetConfigType("json")
	viper.SetConfigFile("config.json")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}

	urls := viper.GetStringSlice("urls")

	for _, url := range urls {
		fmt.Println(url)
		google.GetResult(url)
	}
}

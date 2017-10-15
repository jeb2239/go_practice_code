package main

import (
	"fmt"
	"github.com/spf13/viper"
	"bytes"
)





func initconf(){
	viper.SetConfigName("myconfig") //no file extension
	viper.AddConfigPath("$CONFIG_HOME/config") // allows using env variables
	err := viper.ReadInConfig() // find and read conf
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	
	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
	Hacker: true
	name: steve
	hobbies:
	- skateboarding
	- snowboarding
	- go
	clothing:
	  jacket: leather
	  trousers: denim
	age: 35
	eyes : brown
	beard: true
	`)
	
	viper.ReadConfig(bytes.NewBuffer(yamlExample))
	
	viper.Get("name") // this would be "steve"


}
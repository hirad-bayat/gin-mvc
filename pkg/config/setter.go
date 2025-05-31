package config

//responsible for read the config from yml file

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func Set() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Error while reading the configs")
			log.Fatal("Error while reading the configs")

		} else {
			// Config file was found but another error was produced
			fmt.Println("config file loaded successfully...")
		}
	}

	err := viper.Unmarshal(&configurations)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
		log.Fatal("unable to decode into struct")
	}
	// Config file found and successfully parsed
}

// main.go
package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Name     string `json:"name"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"database"`
}

func main() {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %s", err)
	}

	fmt.Printf("Server running on %s:%d\n", config.Server.Host, config.Server.Port)
	fmt.Printf("Connected to database %s on %s:%d\n", 
		config.Database.Name, 
		config.Database.Host, 
		config.Database.Port)
}
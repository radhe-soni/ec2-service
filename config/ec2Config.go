package config

import (
	"fmt"
	"log"
	"os"

	"github.com/radhe-soni/ec2-service/service"
	"github.com/spf13/viper"
)

// InitConfig to initialize viper and services
func InitConfig() {
	intializeViper()

	viper.Set("awsSession", GetSession())
	viper.Set("ec2Service", service.NewEC2Service())
}
func intializeViper() {
	configPath := getConfigPath()
	log.Printf("loading config from [%s/config.yml] \n", configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
func getConfigPath() string {
	configPath := "./resources"
	if len(os.Args) > 1 {
		cmdArgs := os.Args[1:]
		if cmdArgs[0] != "" {
			configPath = cmdArgs[0]
		}
	}
	return configPath
}

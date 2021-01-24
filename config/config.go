package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	host string `yaml:"host"`
	port int    `yaml:"port"`
}

func (c *Config) GetConfig(configPath string) {
	_, err := os.Stat(configPath)
	if err !=nil{
		log.Print("Config file does not exist! Please check")
		os.Exit(1)
	}

	viper.SetConfigFile(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("%v\n", err)
	}
	if viper.IsSet("service.host") {
		c.host = viper.Get("service.host").(string)
	} else {
		log.Printf("Service ip not set! Please Check")
		os.Exit(1)
	}
	if viper.IsSet("service.port") {
		c.port = viper.Get("service.port").(int)
	} else {
		log.Printf("Service port not set! Please Check")
		os.Exit(1)
	}
}

func (c *Config) GetPort() int {
	return c.port
}
func (c *Config) GetHost() string {
	return c.host
}

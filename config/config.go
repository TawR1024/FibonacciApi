package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	host string `yaml:"host"`
	port int    `yaml:"port"`
}

func (c *Config) GetConfig(configPath string) {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("%v\n", err)
	}
	if viper.IsSet("service.host") {
		c.host = viper.Get("service.host").(string)
	} else {
		log.Printf("Service ip not set! Please Check")
		return
	}
	if viper.IsSet("service.port") {
		c.port = viper.Get("service.port").(int)
	} else {
		log.Printf("Service port not set! Please Check")
		return
	}
}

func (c *Config) GetPort() int {
	return c.port
}
func (c *Config) GetHost() string {
	return c.host
}

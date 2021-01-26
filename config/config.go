package config

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

type Constants struct {
	host      string `yaml:"host"`
	port      int    `yaml:"port"`
	redisPort int    `yaml:"port"`
	redisIP   string `yaml:"ip"`
	redisDB   int    `yaml:"db"`
	redisPass string `yaml:"password"`
}

type Config struct {
	Constants
	RedisClient *redis.Client
}

func New(configPath *string) (*Config, error) {
	config := Config{}
	constants := ReadConfig(*configPath)
	if constants == nil {
		return &config, errors.New("Config Read Error. Please check config file.")
	}
	config.Constants = *constants
	config.RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.redisIP + ":" + strconv.Itoa(config.redisPort),
		Password: config.redisPass, // no password set
		DB:       config.redisDB,   // use default DB
	})
	return &config, nil
}

func ReadConfig(configPath string) *Constants {
	var constants Constants
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("%v\n", err)
	}
	if viper.IsSet("service.host") {
		constants.host = viper.Get("service.host").(string)
	} else {
		log.Printf("Service ip not set! Please Check")
		return nil
	}
	if viper.IsSet("service.port") {
		constants.port = viper.Get("service.port").(int)
	} else {
		log.Printf("Service port not set! Please Check")
		return nil
	}
	if viper.IsSet("redis.port") {
		constants.redisPort = viper.Get("redis.port").(int)
	} else {
		log.Printf("redis port")
		return nil
	}
	if viper.IsSet("redis.ip") {
		constants.redisIP = viper.Get("redis.ip").(string)
	} else {
		log.Printf("redis ip")
		return nil
	}
	if viper.IsSet("redis.db") {
		constants.redisDB = viper.Get("redis.db").(int)
	} else {
		log.Printf("redis db")
		return nil
	}
	if viper.IsSet("redis.pass") {
		constants.redisPass = viper.Get("redis.pass").(string)
		log.Print(viper.Get("redis.pass").(string))
	} else {
		constants.redisPass = ""
		//return nil
	}
	return &constants
}

func (c *Constants) GetPort() int {
	return c.port
}
func (c *Constants) GetHost() string {
	return c.host
}

func (c *Constants) GetRedisIp() string {
	return c.redisIP
}
func (c *Constants) GetRedisPort() int {
	return c.redisPort
}
func (c *Constants) GetRedisDB() int {
	return c.redisDB
}
func (c *Constants) GetRedisPass() string {
	return c.redisPass
}

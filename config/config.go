package config

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"golang.org/x/net/context"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
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
		Password: config.redisPass,
		DB:       config.redisDB,
	})
	status := checkRedisConnection(config.RedisClient)
	if status != nil {
		os.Exit(1) //Stop running while redis not connected
	}
	return &config, nil
}

func checkRedisConnection(client *redis.Client) error {
	//Check connection to Redis during 3 seconds
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel()

	ans, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Conection to redis failed! %v", err)
		return err
	}
	log.Printf("Redis connection OK, %v", ans)
	return nil
}

// ReadConfig get configPath, read config
// return nil if necessary variables not set
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
		constants.redisPass = "" // using empty password if password not set
	}
	return &constants
}

func (c *Constants) GetPort() int {
	return c.port
}
func (c *Constants) GetHost() string {
	return c.host
}

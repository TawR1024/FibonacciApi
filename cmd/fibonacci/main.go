package main

import (
	"flag"
	"github.com/TawR1024/FibonacciApi/calculator"
	"github.com/TawR1024/FibonacciApi/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var conf *config.Config

func main() {

	var configFilePath = flag.String("config", "configFile", "Setting the configuration file")
	flag.Parse()
	_, err := os.Stat(*configFilePath)
	if err == nil {
		log.Printf("Using custom config path: %s", *configFilePath)
		conf, _ = config.New(configFilePath)
	} else {

		defaultPath := "/etc/fibonacci/config.yaml" // load config from default config path
		log.Printf("Using default config path: %s", defaultPath)
		conf, _ = config.New(&defaultPath)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(5 * time.Second))

	router.Get("/fibonacci_small", calculator.New(conf).CountFibonacciBinet)
	router.Get("/fibonacci_big", calculator.New(conf).CountFibonacciRecursive)

	log.Fatal(http.ListenAndServe(conf.GetHost()+":"+strconv.Itoa(conf.GetPort()), router))
}

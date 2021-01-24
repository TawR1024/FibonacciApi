package main

import (
	"github.com/TawR1024/FibonacciApi/calculator"
	"github.com/TawR1024/FibonacciApi/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"strconv"
	"time"
)

var conf config.Config

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(5 * time.Second))

	router.Get("/fibonacci_small", calculator.CountFibonacciBinet)
	router.Get("/fibonacci_big", calculator.CountFibonacciRecurcive)

	http.ListenAndServe(conf.GetHost()+":"+strconv.Itoa(conf.GetPort()), router)
}

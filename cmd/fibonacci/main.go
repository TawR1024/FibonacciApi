package fibonacci

import (
	"github.com/TawR1024/FibonacciApi/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"strconv"
	"time"
)

var conf config.Config

func main(){
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(5 * time.Second))

	router.Get("/fibonacci", func(writer http.ResponseWriter, request *http.Request) {
		
	})


	http.ListenAndServe(conf.GetHost()+":"+strconv.Itoa(conf.GetPort()), router)
}
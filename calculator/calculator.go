package calculator

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"sync"

	"github.com/TawR1024/FibonacciApi/config"
	"github.com/TawR1024/FibonacciApi/connector"
)

var sqrtOfFive = math.Sqrt(5)

type Body struct {
	From int `json:"from"`
	To   int `json:"to"`
}
type Config struct {
	*config.Config
}

func New(configuration *config.Config) *Config {
	return &Config{configuration}
}
func (config *Config) CountFibonacciBinet(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Range is empty!", 400)
		return
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print("Request body read error")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("Can`t read request body"))
		if err != nil {
			log.Printf("Response error %v", err)
		}
	}

	data := &Body{}
	err = json.Unmarshal(body, data)
	if err != nil {
		log.Printf("Can`t unmarshal request body in CountFibonacciBinet() %v", err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Bad params! to must be greater than from, both vars must be positive"))
		if err != nil {
			log.Printf("Response error %v", err)
		}
		return
	}
	if data.From > data.To || data.From < 0 || data.To < 0 {
		log.Printf("Fishy arguments from: %d to: %d", data.From, data.To)

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("Bad params! to must be greater than from, both vars must be positive"))
		if err != nil {
			log.Printf("Response error %v", err)
		}
		return
	}

	var wg sync.WaitGroup
	dataChannel := make(chan map[int]float64, data.To-data.From+1)
	for i := data.From; i <= data.To; i++ {
		wg.Add(1)
		go fibonacci(&wg, &dataChannel, i)
	}
	wg.Wait()
	close(dataChannel)
	response := make(map[int]float64)
	for i := range dataChannel {
		for k, v := range i {
			response[k] = v
		}
	}

	payload, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(payload)
	if err != nil {
		log.Printf("Response error %v", err)
	}
	return
}

func fibonacci(w *sync.WaitGroup, c *chan map[int]float64, n int) {
	defer w.Done()
	a := map[int]float64{n: binet(n)}
	*c <- a
}

// Binet Formula
func binet(n int) float64 {
	f := (math.Pow((1+sqrtOfFive)/2, float64(n)) - math.Pow((1-sqrtOfFive)/2, float64(n))) / sqrtOfFive
	return math.Round(f)
}

func (config *Config) CountFibonacciRecursive(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	if r.Body == nil {
		http.Error(w, "Range is empty!", 400)
		return
	}
	defer r.Body.Close()

	data := &Body{}
	err := json.Unmarshal(body, data)
	if err != nil {
		log.Printf("Can`t unmarshal request body in CountFibonacciRecursive() %v", err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Bad params! to must be greater than from, both vars must be positive"))
		if err != nil {
			log.Printf("Response error %v", err)
		}
		return
	}
	if data.From > data.To || data.From < 0 || data.To < 0 {
		log.Printf("Fishy arguments from: %d to: %d", data.From, data.To)

		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Bad params! to must be greater than from, both vars must be positive"))
		if err != nil {
			log.Printf("Response error %v", err)
		}
		return
	}
	response := make(map[int]string)

	for i := data.From; i <= data.To; i++ {
		response[i] = config.fibonacciBig(i).String()
	}
	log.Print(response)

	payload, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(payload)
	if err != nil {
		log.Printf("Response error %v", err)
	}
	return
}

// Recursive fibonacci formula
func (config *Config) fibonacciBig(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	if n == 1 || n == 2 {
		return big.NewInt(1)
	}
	ifCached := connector.GetBigKey(config.RedisClient, n) // check if cached
	if ifCached == nil {
		prev := config.fibonacciBig(n - 1)
		prev2 := config.fibonacciBig(n - 2)
		prev.Add(prev, prev2)
		connector.SetBigKey(config.RedisClient, n, *prev)
		return prev
	} else {
		return ifCached
	}
}

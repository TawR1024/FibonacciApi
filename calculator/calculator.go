package calculator

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"sync"
)

var sqrtOfFive = math.Sqrt(5)

type Body struct {
	From int `json:"from"`
	To   int `json:"to"`
}

func CountFibonacciBinet(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Print("Request body read error")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("Can`t read request body"))
		if err != nil {
			log.Printf("Response error %v", err)
		}
	}

	if r.Body == nil {
		http.Error(w, "Range is empty!", 400)
		return
	}
	data := &Body{}
	err = json.Unmarshal(body, data)
	if err !=nil{

	}
	if data.From > data.To || data.From < 0 || data.To < 0 {
		log.Printf("Fishy arguments from: %d to: %d", data.From, data.To)

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("Bad params! to must be greater than from, both wars must be positive"))
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
	log.Print("Send Data to user")
	w.Write(payload)
}

func fibonacci(w *sync.WaitGroup, c *chan map[int]float64, n int) {
	defer w.Done()
	a := map[int]float64{n: binet(n)}
	*c <- a
}

func binet(n int) float64 {
	f := (math.Pow((1+sqrtOfFive)/2, float64(n)) - math.Pow((1-sqrtOfFive)/2, float64(n))) / sqrtOfFive
	return math.Round(f)
}

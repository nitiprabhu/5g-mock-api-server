package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	r   = mux.NewRouter()
	adr = os.Getenv("ADDRESS")
)

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	var p map[string]string

	err := json.NewDecoder(request.Body).Decode(&p)
	if err != nil {
		log.Println(err.Error(), http.StatusBadRequest)
		responseWriter.Write([]byte("true"))
		return
	}
	log.Println("Received:", p)

	resp, err := http.Get(adr + "/predict")

	if err != nil {
		log.Println(err)
		responseWriter.Write([]byte("true"))
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		responseWriter.Write([]byte("true"))
		return
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
	responseWriter.Write([]byte(sb))
	return
}

func main() {
	log.Println("Address:", adr)

	r.HandleFunc("/products", handler).Methods("POST")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

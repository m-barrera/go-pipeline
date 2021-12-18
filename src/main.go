package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HelloWorld).Methods("GET")
	err := http.ListenAndServe(":80", router)
	if err != nil {
		panic("Failed serve")
	}
}

type Response struct {
	Text string `json:"text"`
}

func HelloWordTxt() Response {
	return Response{Text: "HelloWorld"}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(HelloWordTxt())
	if err != nil {
		panic("failed encode")
	}
}

package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HelloWorld).Methods("GET")
	http.ListenAndServe(":80", router)
}

type Response struct {
	Text string `json:"text"`
}

func HelloWordTxt() Response {
	return Response{Text: "HelloWorld"}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(HelloWordTxt())
}

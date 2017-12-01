package main

import (
	"encoding/json"
	"net/http"
//	"github.com/gorilla/handlers"
//	"github.com/gorilla/mux"

)


type Dogpark struct {
Name   string `json:"name"`
Address string `json:"address"`
Id     string `json:"id"`
}

func main() {
//	r := mux.NewRouter()
//	r.HandleFunc("/dogparks", foo)
	http.HandleFunc("/dogparks", foo)
//	http.ListenAndServe(":3000", handlers.CORS()(r))
	http.ListenAndServe(":3000", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := []Dogpark{{"Twin Lakes Dog Park", "Dog Park Street", "1"},{"Boulder Dog Park", "Boulder Street", "2"}}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}


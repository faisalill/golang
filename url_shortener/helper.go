package main

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
}

var Store = map[int]string{}

func Server(address string) *ApiServer {
	return &ApiServer{
		addr: address,
	}
}

func (s *ApiServer) Run() {
	mux := mux.NewRouter()

	mux.HandleFunc("/{id}", homeHandler).Methods("GET")
	mux.HandleFunc("/health", healthHandler).Methods("GET")
	mux.HandleFunc("/register", registerHandler).Methods("POST")

	http.ListenAndServe(s.addr, mux)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for key, val := range Store {
		if key == id {
			http.Redirect(w, r, val, http.StatusPermanentRedirect)
		} else {
			http.Error(w, "Invalid URL id", http.StatusBadRequest)
		}
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var val map[string]string
	err := json.NewDecoder(r.Body).Decode(&val)
	checkErr(err)

	if _, exists := val["url"]; exists {

		log.Println("URL: ", val["url"])

		key := rand.IntN(1000)

		Store[key] = val["url"]

		msg := "Shortened Url: http://localhost:8000/" + strconv.Itoa(key)
		w.Write([]byte(msg))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing url from the payload"))
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

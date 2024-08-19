package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, val any) error {
	w.Header().Set("Content-Type", "application/json;")
	return json.NewEncoder(w).Encode(val)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if err := f(w, req); err != nil {
			WriteJSON(w, http.StatusBadRequest, err.Error())
		}
	}
}

type ApiServer struct {
	serverAddress string
}

func NewApiServer(address string) *ApiServer {
	return &ApiServer{
		serverAddress: address,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.HandleFunc("/account", makeHttpHandleFunc(s.handleAccount))

	log.Println("Server Starting at: ", s.serverAddress)

	log.Fatal(http.ListenAndServe(s.serverAddress, subRouter))
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, req *http.Request) error {
	if req.Method == "GET" {
		return s.handleGetAccount(w, req)
	}
	if req.Method == "POST" {
		return s.handleCreateAccount(w, req)
	}

	if req.Method == "DELETE" {
		return s.handleDeleteAccount(w, req)
	}

	return fmt.Errorf("Method not allowed : %s", req.Method)
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, req *http.Request) error {
	account := NewAccount("Miya", "Bhai")
	WriteJSON(w, http.StatusOK, account)
	return nil
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, req *http.Request) error {
	return nil
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, req *http.Request) error {
	return nil
}

func (s *ApiServer) handleTransfer(w http.ResponseWriter, req *http.Request) error {
	return nil
}

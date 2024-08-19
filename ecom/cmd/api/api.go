package api

import (
	"database/sql"
	"ecom/service/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	Address string
	db      *sql.DB
}

func NewServer(listeningAddress string, db *sql.DB) *ApiServer {
	return &ApiServer{
		Address: listeningAddress,
		db:      db,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userService := user.NewHandler()
	userService.RegisterRoutes(subRouter)

	log.Fatal(http.ListenAndServe(s.Address, subRouter))
}

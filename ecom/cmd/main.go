package main

import (
	"database/sql"
	"ecom/cmd/api"
	"ecom/config"
	"ecom/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySqlStorage(
		mysql.Config{
			User:   config.Envs.DBUser,
			Passwd: config.Envs.DBPassword,
			Addr:   config.Envs.DBAddress,
			DBName: config.Envs.DBName,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	initDatabase(db)

	server := api.NewServer(":8000", db)

	log.Println("Server Started at: ", server.Address)
	server.Run()
}

func initDatabase(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Successfully Connected")
}

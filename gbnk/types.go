package main

import "math/rand"

type Account struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AccNumber int    `json:"accNumber"`
	Balance   int    `json:"balance"`
}

func NewAccount(firstname string, lastname string) *Account {
	return &Account{
		Id:        rand.Intn(10000),
		AccNumber: rand.Intn(1000000),
		FirstName: firstname,
		LastName:  lastname,
	}
}

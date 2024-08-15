package main

import "net/http"

func homeHandler(res http.ResponseWriter, req *http.Request) {
	message := "Yuar Path: " + req.URL.Path

	res.Write([]byte(message))
}

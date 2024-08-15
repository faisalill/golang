package main

import (
	"html/template"
	"log"
	"net/http"
)

func homeHandler(res http.ResponseWriter, req *http.Request) {
	message := "Yuar Path: " + req.URL.Path

	res.Write([]byte(message))
}

func htmlHandler(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/post.html")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	data := GetPosts()

	err = template.Execute(res, data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

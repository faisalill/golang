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

	data := struct {
		UserId int32
		Id     int32
		Title  string
		Body   string
	}{
		UserId: 1,
		Id:     3,
		Title:  "Title of the Blog",
		Body:   "Body of the whole blog. Body of the whole blog.",
	}

	err = template.Execute(res, data)
	if err != nil {
		http.Error(res, err.Error(), 402)
	}
}

package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	UserId int32
	Id     int32
	Title  string
	Body   string
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	postsUrl := "https://jsonplaceholder.typicode.com/posts"

	resp, err := http.Get(postsUrl)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var Posts []Post

	err = json.Unmarshal(data, &Posts)
	if err != nil {
		panic(err)
	}

	temp, err := template.ParseFiles("./post.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(res, Posts)
}

func main() {
	http.HandleFunc("/posts", homeHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

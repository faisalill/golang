package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var PostsUrl string = "https://jsonplaceholder.typicode.com/posts"

type Post struct {
	UserId int32
	Id     int32
	Title  string
	Body   string
}

func GetPosts() {
	resp, err := http.Get(PostsUrl)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var Posts []Post

	err = json.Unmarshal(data, &Posts)

	for _, value := range Posts {
		fmt.Println("...")
		fmt.Println("User Id", value.UserId)
		fmt.Println("Id", value.Id)
		fmt.Println("Title", value.Title)
		fmt.Println("Body", value.Body)
		fmt.Println("...")
	}
}

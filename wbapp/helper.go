package main

import (
	"encoding/json"
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

func GetPosts() []Post {
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

	return Posts
}

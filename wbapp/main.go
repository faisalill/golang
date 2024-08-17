package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Post struct {
	id     int32
	userid int32
	title  string
	body   string
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getPosts(db *sql.DB) []Post {
	rows, err := db.Query("SELECT * FROM posts")
	checkErr(err)

	var Posts []Post

	for rows.Next() {
		var id int32
		var userid int32
		var title string
		var body string

		err = rows.Scan(&id, &userid, &title, &body)
		checkErr(err)

		Posts = append(Posts, Post{id, userid, title, body})

	}

	return Posts
}

func uploadPost(db *sql.DB, userid int32, title string, body string) {
	_, err := db.Exec("INSERT INTO posts(userid, title, body) values($1, $2, $3)", userid, title, body)
	checkErr(err)
	fmt.Println("Succesfully Uploaded Post")
}

func deletePost(db *sql.DB, id int32, userid int32) {
	_, err := db.Exec("DELETE FROM posts WHERE id = $1 AND userid = $2", id, userid)
	checkErr(err)
	fmt.Println("Succesfully Deleted Post")
}

func main() {
	db, err := sql.Open("postgres", "user=postgres password=password dbname=wbapp sslmode=disable")
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	fmt.Println("Connected to Database Successfully")

	posts := getPosts(db)

	fmt.Println(posts)
}

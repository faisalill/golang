package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Post struct {
	Id     int32
	Userid int32
	Title  string
	Body   string
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

func showPostsHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts := getPosts(db)

		temp, err := template.ParseFiles("posts.html")
		checkErr(err)

		var buf bytes.Buffer
		err = temp.ExecuteTemplate(&buf, "posts.html", posts)
		checkErr(err)

		c.Set("Content-Type", "text/html; charset=utf-8")
		return c.SendString(buf.String())
	}
}

func formHandler(c *fiber.Ctx) error {
	temp, err := template.ParseFiles("form.html")
	checkErr(err)

	var buf bytes.Buffer
	err = temp.ExecuteTemplate(&buf, "form.html", nil)

	c.Set("Content-Type", "text/html; charset=utf-8")

	return c.SendString(buf.String())
}

func main() {
	db, err := sql.Open("postgres", "user=postgres password=password dbname=wbapp sslmode=disable")
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	fmt.Println("Connected to Database Successfully")
	fmt.Println("Starting Server")

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		if c.Path() != "/posts" && c.Path() != "/new" && c.Path() != "/createpost" {
			return c.Redirect("/posts")
		}

		return c.Next()
	})

	app.Get("/posts", showPostsHandler(db))
	app.Get("/new", formHandler)
	app.Post("/createpost", func(c *fiber.Ctx) error {
		title := c.FormValue("title")
		body := c.FormValue("body")

		fmt.Printf("Title: %s Body: %s", title, body)

		c.Set("Content-Type", "text/html; charset=utf-8")
		return c.SendString("<h1>Post Created Successfully</h1>")
	})

	log.Fatal(app.Listen(":3000"))
}

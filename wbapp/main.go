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

func createUser(db *sql.DB, username string, password string) {
	_, err := db.Query("INSERT INTO users(name, password) values($1, $2)", username, password)
	checkErr(err)
	fmt.Println("Successfully Created New User")
}

func checkUser(db *sql.DB, username string, password string) bool {
	var count int
	err := db.QueryRow(
		"SELECT COUNT(*) FROM users WHERE name = $1 AND password = $2", username, password).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
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
	checkErr(err)

	c.Set("Content-Type", "text/html; charset=utf-8")

	return c.SendString(buf.String())
}

func registrationHandler(c *fiber.Ctx) error {
	temp, err := template.ParseFiles("register.html")
	checkErr(err)

	var buf bytes.Buffer
	err = temp.ExecuteTemplate(&buf, "register.html", nil)
	checkErr(err)

	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(buf.String())
}

func registerUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		createUser(db, username, password)

		return c.SendString("User Created Successfully")
	}
}

func loginHandler(c *fiber.Ctx) error {
	temp, err := template.ParseFiles("login.html")
	checkErr(err)

	var buf bytes.Buffer
	err = temp.ExecuteTemplate(&buf, "login.html", nil)
	checkErr(err)

	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(buf.String())
}

func loginSessionHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		isUserRegistered := checkUser(db, username, password)

		if isUserRegistered {
			return c.SendString("Session Created ")
		} else {
			return c.SendString("Wrong Credentials")
		}
	}
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
		if c.Path() != "/posts" && c.Path() != "/new" && c.Path() != "/createpost" && c.Path() != "/register" && c.Path() != "/login" {
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

	app.Get("/register", registrationHandler)
	app.Post("/register", registerUserHandler(db))
	app.Get("/login", loginHandler)
	app.Post("/login", loginSessionHandler(db))

	log.Fatal(app.Listen(":3000"))
}

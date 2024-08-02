package main

import (
	"log"
	"net/http"
)

func homeHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}

	response.Write([]byte("THIS IS WORKING"))
}

func showSnippetHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Show Snippets"))
}

func createSnippetHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		response.Header().Set("Allow", "POST")
		http.Error(response, "Method Not Allowed", 405)
		return
	}
	response.Write([]byte("Create Snippets"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet", showSnippetHandler)
	mux.HandleFunc("/snippet/create", createSnippetHandler)

	log.Println("Starting server")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}

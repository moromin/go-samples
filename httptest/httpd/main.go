package main

import (
	"fmt"
	"go-samples/httptest/httpd/handler"
	"go-samples/httptest/platform/newsfeed"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := ":3000"
	feed := newsfeed.New()
	feed.Add(newsfeed.Item{
		Title: "Hello",
		Post:  "World",
	})

	r := chi.NewRouter()

	// Get newsfeed
	r.Get("/newsfeed", handler.NewsfeedGet(feed))

	// Post newsfeed
	r.Post("/newsfeed", handler.NewsfeedPost(feed))

	fmt.Println("Serving on port ", port)
	http.ListenAndServe(port, r)
}

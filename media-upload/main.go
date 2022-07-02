package main

import (
	"fmt"
	"go-samples/media-upload/file"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file.UploadHandler(w, r)
}

func setupRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("FILE UPLOAD MONITOR")

	setupRoutes()
}

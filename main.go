package main

import (
	"net/http"
	"url-shortener/handlers"
	"url-shortener/storage"
)

func main() {
	store := storage.NewStorage()
	handler := handlers.NewHandler(store)

	http.HandleFunc("/shorten", handler.ShortenURL)
	http.HandleFunc("/", handler.RedirectURL)

	http.ListenAndServe(":8080", nil)
}
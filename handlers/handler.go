package handlers

import (
    "fmt"
    "net/http"
    "url-shortener/models"
    "url-shortener/storage"
)

type Handler struct {
    storage *storage.Storage
}

func NewHandler(storage *storage.Storage) *Handler {
    return &Handler{
        storage: storage,
    }
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
        return
    }

    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid form", http.StatusBadRequest)
        return
    }

    originalURL := r.FormValue("url")
    if originalURL == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }

    shortURL := models.GenerateShortURL()
    h.storage.Save(shortURL, originalURL)

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Shortened URL: %s\n", shortURL)
}

// ...existing code...
func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodGet {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    shortURL := r.URL.Path[1:]
    fmt.Println("shortURL:", shortURL)

    originalURL, exists := h.storage.Get(shortURL)
    if !exists {
        http.Error(w, "URL Not Found", http.StatusNotFound)
        return
    }

    http.Redirect(w, r, originalURL, http.StatusSeeOther)
}
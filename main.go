package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type URLStore struct {
	sync.RWMutex
	urls     map[string]string
	analytics map[string]int
}

func NewURLStore() *URLStore {
	return &URLStore{
		urls:      make(map[string]string),
		analytics: make(map[string]int),
	}
}

func (s *URLStore) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var request struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	shortCode := generateShortCode()
	s.Lock()
	s.urls[shortCode] = request.URL
	s.Unlock()

	response := map[string]string{"short_url": fmt.Sprintf("http://localhost:8080/%s", shortCode)}
	json.NewEncoder(w).Encode(response)
}

func (s *URLStore) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	s.RLock()
	originalURL, exists := s.urls[shortCode]
	s.RUnlock()

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	s.Lock()
	s.analytics[shortCode]++
	s.Unlock()

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func (s *URLStore) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	s.RLock()
	clicks, exists := s.analytics[shortCode]
	s.RUnlock()

	if !exists {
		http.Error(w, "No analytics data found for this URL", http.StatusNotFound)
		return
	}

	response := map[string]int{"clicks": clicks}
	json.NewEncoder(w).Encode(response)
}

func generateShortCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	store := NewURLStore()
	r := mux.NewRouter()

	r.HandleFunc("/shorten", store.ShortenURL).Methods("POST")
	r.HandleFunc("/{shortCode}", store.Redirect).Methods("GET")
	r.HandleFunc("/analytics/{shortCode}", store.GetAnalytics).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

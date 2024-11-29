package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Struktur untuk respons JSON
type Response struct {
	Message string `json:"message"`
}

// Handler untuk endpoint root "/"
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Buat respons JSON
	response := Response{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

// Handler untuk endpoint "/greet/{name}"
func greetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Ekstrak nama dari path
	path := strings.TrimPrefix(r.URL.Path, "/greet/")
	name := strings.TrimSpace(path)

	// Jika nama kosong, beri pesan default
	if name == "" {
		name = "Guest"
	}

	// Buat respons JSON
	response := Response{Message: fmt.Sprintf("Hello, %s!", name)}
	json.NewEncoder(w).Encode(response)
}

// Middleware sederhana untuk logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Mendaftarkan handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorldHandler)
	mux.HandleFunc("/greet/", greetHandler)

	// Tambahkan middleware untuk logging
	loggedMux := loggingMiddleware(mux)

	// Jalankan server di port 8080
	fmt.Println("Server berjalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		fmt.Println("Error menjalankan server:", err)
	}
}

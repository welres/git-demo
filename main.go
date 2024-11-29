package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, `{"message": "Hello, World!"}`)
}

func main() {
	http.HandleFunc("/", helloWorldHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error menjalankan server:", err)
	}
}

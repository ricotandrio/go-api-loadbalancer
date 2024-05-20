package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"cloud-api/models"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "cloud-api. This is the root path\n\n")
	fmt.Fprintln(w, "localhost:8080/hello")
}

func helloMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	message := models.Hello{
		Message: "Hello, World!",
	}

	json.NewEncoder(w).Encode(message)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/info", helloMessageHandler)

	addr := ":8080"

	fmt.Printf("Server available at http://localhost%s\n", addr)

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		panic(err)
	}
}

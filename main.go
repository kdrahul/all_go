package main

import (
	"fmt"
	"log"
	"net/http"
)

type Product struct {
  Name string `json:"name"`
  Price string `json:"price"`
}

type Products []Product

type productHandler struct {
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello")
}

func main() {
  port := ":8080"
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(port, nil))
}

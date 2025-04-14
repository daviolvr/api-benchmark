package main

import (
	"benchmark-api/internal/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/benchmark", api.BenchmarkHandler)
	http.HandleFunc("/results", api.ResultsHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bem-vindo à API de benchmark!")
}

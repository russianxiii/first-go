package main

import (
	"fmt"
	"net/http"

	"first-go/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)

	http.HandleFunc("/sum", handlers.SumHandler)

	http.HandleFunc("/sort", handlers.SortHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"

	"first-go/handlers"
)

func main() {
	// Обработка запроса на /hello
	http.HandleFunc("/hello", handlers.HelloHandler)

	http.HandleFunc("/sum", handlers.SumHandler)

	// Запуск сервера на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// Обработка запроса на /hello
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		// Получаем параметры из query
		aStr := r.URL.Query().Get("a")
		bStr := r.URL.Query().Get("b")

		// Преобразуем строки в числа
		a, errA := strconv.Atoi(aStr)
		b, errB := strconv.Atoi(bStr)

		// Обработка ошибок
		if errA != nil || errB != nil {
			http.Error(w, "Оба параметра должны быть целыми числами (например, ?a=3&b=4)", http.StatusBadRequest)
			return
		}

		// Считаем сумму
		sum := a + b

		// Возвращаем результат
		fmt.Fprintf(w, "Результат: %d", sum)
	})

	// Запуск сервера на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

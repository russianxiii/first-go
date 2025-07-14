package handlers

import "net/http"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Отправляем ответ "Hello, world!" на запрос к /hello
	w.Write([]byte("Hello, world!"))
}

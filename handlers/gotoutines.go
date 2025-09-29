package handlers

import (
	"net/http"
	"sync"
	"time"
)

func sayHello(w http.ResponseWriter, wg *sync.WaitGroup, message string) {
	defer wg.Done()

	w.Write([]byte(message))
}

func RoutinesHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	w.Write([]byte("Goroutines!\n\n"))

	arr := [4]string{
		"User 1: hello!\n",
		"User 2: hi!\n",
		"User 1: how are you?\n",
		"User 2: fine!\n",
	}

	for _, value := range arr {
		wg.Add(1)
		go sayHello(w, &wg, value)
		time.Sleep(time.Millisecond)
	}

	wg.Wait()
}

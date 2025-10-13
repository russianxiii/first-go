package handlers

import (
	"net/http"
	"sync"

	"fmt"
)

func sayHello(wg *sync.WaitGroup, mu *sync.Mutex, message string, messages *[]string, idx int) {
	defer wg.Done()

	mu.Lock()
	(*messages)[idx] = message
	mu.Unlock()
}

func RoutinesHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	w.Write([]byte("Goroutines!\n\n"))

	arr := []string{
		"\n",
		"User 1: hello!\n",
		"User 2: hi!\n",
		"User 1: how are you?\n",
		"User 2: fine!\n",
	}

	messages := make([]string, len(arr))

	for i, value := range arr {
		wg.Add(1)
		go sayHello(&wg, &mu, value, &messages, i)
	}

	wg.Wait()

	for _, msg := range messages {
		w.Write([]byte(msg))
		fmt.Println(msg)
	}
}

package handlers

import (
	"log"
	"net/http"
	"sync"

	"fmt"
)

func sayHello(wg *sync.WaitGroup, mu *sync.Mutex, message string, messages *[]string, idx int) {
	defer wg.Done()

	mu.Lock()
	var err error
	if message == "User X: 404" {
		err = simulateWriteError()
	} else {
		(*messages)[idx] = message
	}
	mu.Unlock()

	if err != nil {
		log.Println("Error writing message:", err)
	}
}

func simulateWriteError() error {
	return fmt.Errorf("ошибка записи")
}

func RoutinesHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	w.Write([]byte("Goroutines and mutex!\n\n"))

	arr := []string{
		"User 1: hello!",
		"User 2: hi!",
		"User 1: how are you?",
		"User X: 404",
		"User 2: fine!",
	}

	messages := make([]string, len(arr))

	for i, value := range arr {
		wg.Add(1)
		go sayHello(&wg, &mu, value, &messages, i)
	}

	wg.Wait()

	for idx, msg := range messages {
		if msg == "" {
			continue
		}
		w.Write([]byte(msg + "\n"))
		fmt.Printf("%d %s\n", idx, msg)
	}
}

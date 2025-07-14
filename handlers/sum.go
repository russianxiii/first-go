package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func SumHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, errA := strconv.Atoi(aStr)
	b, errB := strconv.Atoi(bStr)

	if errA != nil || errB != nil {
		http.Error(w, "Оба параметра должны быть целыми числами (например, ?a=3&b=4)", http.StatusBadRequest)
		return
	}

	sum := a + b
	fmt.Fprintf(w, "Результат: %d", sum)
}

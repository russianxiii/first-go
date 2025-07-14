package handlers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func SortHandler(w http.ResponseWriter, r *http.Request) {
	numsStr := r.URL.Query().Get("nums")
	if numsStr == "" {
		http.Error(w, "Параметр nums обязателен, например ?nums=3,2,1", http.StatusBadRequest)
		return
	}

	// Разбиваем строку по запятым
	numsSlice := strings.Split(numsStr, ",")

	// Парсим в срез int
	nums := make([]int, 0, len(numsSlice))
	for _, s := range numsSlice {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			http.Error(w, "Все значения должны быть целыми числами", http.StatusBadRequest)
			return
		}
		nums = append(nums, n)
	}

	// Сортируем срез
	sort.Ints(nums)

	// Формируем строку результата
	strNums := make([]string, len(nums))
	for i, n := range nums {
		strNums[i] = strconv.Itoa(n)
	}
	result := strings.Join(strNums, ",")

	// Возвращаем
	fmt.Fprintf(w, "Отсортировано: %s", result)
}

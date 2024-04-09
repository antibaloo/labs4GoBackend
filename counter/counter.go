package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http"
	"strconv" // вдруг понадобиться вам ;)
)

var counter int

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, counter)
	}
	if r.Method == "POST" {
		r.ParseForm()
		numberString := r.FormValue("count")

		num, err := strconv.ParseInt(numberString, 10, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
			return
		}

		counter = counter + int(num)

	}
}
func main() {
	counter = 0
	http.HandleFunc("/count", handler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

package main

import (
	"net/http"
)

// функция main вызывается автоматически при запуске приложения
func main() {
	if err := run(); err != nil {
		panic(err)
	}

}

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func run() error {

	mux := http.NewServeMux()
	mux.HandleFunc(`/`, POSThook)
	mux.HandleFunc(`/EwHXdJfB`, GEThook)
	return http.ListenAndServe(`:8080`, mux)
}

var URLtoken map[string]string = map[string]string{

	"shortener": "http://localhost:8080/EwHXdJfB",
	"full":      "https://practicum.yandex.ru/",
}

func POSThook(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost || r.Header.Get("Content-Type") != "text/plain" {
		// проверка на POST-запросы и на заголовок
		http.Error(w, "", 400)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(201)
	w.Write([]byte(URLtoken["shortener"]))

}

func GEThook(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		// проверка на GET-запросы
		http.Error(w, "", 400)
		return
	}

	w.WriteHeader(307)
	w.Write([]byte(URLtoken["full"]))

}

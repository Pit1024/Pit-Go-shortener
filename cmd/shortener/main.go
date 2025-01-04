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

	return http.ListenAndServe(`:8080`, http.HandlerFunc(webhook))
}

var URLtoken map[string]string = map[string]string{

	"shortener": "http://localhost:8080/EwHXdJfB",
	"full":      "https://practicum.yandex.ru/",
}

func webhook(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost && r.Header.Get("Content-Type") == "text/plain" {
		// проверка на POST-запросы

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(201)
		w.Write([]byte(URLtoken["shortener"]))
	} else if r.Method == http.MethodGet {

		w.WriteHeader(307)
		w.Write([]byte(URLtoken["full"]))
	} else {

		http.Error(w, "", 400)
	}

}

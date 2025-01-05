package main

import (
	"net"
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

var URLmap map[string][]string

// функция webhook — обработчик HTTP-запроса
func webhook(w http.ResponseWriter, r *http.Request) {

	// парсим переданный шаблон, проверяем на ошибки
	host, port, err := net.SplitHostPort(r.Host)
	if err != nil {

		http.Error(w, "", 400)
	}

	// сохраняем пришедший адрес
	URLmap = map[string][]string{"shortener": {host, port}}

	// Если запрос POST и нет /{id}, то выдаем хост и порт и статус 201
	if r.Method == http.MethodPost && port == "" {
		http.Redirect(w, r, host, 201)

		// Если запрос GET и есть /{id}, то выдаем обычный хост и статус 201
	} else if r.Method == http.MethodGet && port != "" {
		http.Redirect(w, r, host, 307)

		// в любом другом случае ошибка 400
	} else {

		http.Error(w, "", 400)
	}

}

//OLD
// package main

// import (
// 	"net/http"
// )

// // функция main вызывается автоматически при запуске приложения
// func main() {
// 	if err := run(); err != nil {
// 		panic(err)
// 	}

// }

// // функция run будет полезна при инициализации зависимостей сервера перед запуском
// func run() error {

// 	mux := http.NewServeMux()
// 	mux.HandleFunc(`/`, POSThook)
// 	mux.HandleFunc(`/EwHXdJfB`, GEThook)
// 	return http.ListenAndServe(`:8080`, mux)
// }

// var URLtoken map[string]string = map[string]string{

// 	"shortener": "http://localhost:8080/EwHXdJfB",
// 	"full":      "http://yandex.ru/",
// }

// func POSThook(w http.ResponseWriter, r *http.Request) {

// 	if r.Method != http.MethodPost { //|| r.Header.Get("Content-Type") != "text/plain" {
// 		// проверка на POST-запросы и на заголовок
// 		http.Error(w, "", 400)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "text/plain")
// 	w.WriteHeader(201)
// 	w.Write([]byte(URLtoken["shortener"]))

// }

// func GEThook(w http.ResponseWriter, r *http.Request) {

// 	if r.Method != http.MethodGet {
// 		// проверка на GET-запросы
// 		http.Error(w, "", 400)
// 		return
// 	}

// 	w.Header().Set("Location", URLtoken["full"])
// 	w.WriteHeader(307)
// 	w.Write([]byte(URLtoken["full"]))

// }

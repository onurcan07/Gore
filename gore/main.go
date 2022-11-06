package main

import (
	"log"
	"net/http"

	"github.com/gore/actions/book"
	"github.com/gore/actions/user"
	"github.com/gore/persistence"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	DB := persistence.Init()
	book_handler := book.New(DB)
	user_handler := user.New(DB)

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/books", book_handler.Get).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", book_handler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/books", book_handler.Create).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", book_handler.Update).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", book_handler.Delete).Methods(http.MethodDelete)

	r.HandleFunc("/users", user_handler.Create).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", user_handler.Update).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", user_handler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/users", user_handler.Get).Methods(http.MethodGet)
	r.HandleFunc("/users{id}", user_handler.GetById).Methods(http.MethodGet)

	log.Println("Listening ...")
	http.ListenAndServe(":8000", r)
}

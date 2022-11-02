package main

import (
	"log"
	"net/http"

	"github.com/gore/actions/book"
	"github.com/gore/actions/grocery"
	"github.com/gore/persistence"
	"github.com/gorilla/mux"
)

func main() {
	DB := persistence.Init()
	h := book.New(DB)

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/books", h.CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	r.HandleFunc("/allgroceries", grocery.AllGroceries).Methods(http.MethodGet)
	r.HandleFunc("/groceries/{name}", grocery.SingleGrocery).Methods(http.MethodGet)
	r.HandleFunc("/groceries", grocery.GroceriesToBuy).Methods(http.MethodPost)
	r.HandleFunc("/groceries/{name}", grocery.UpdateGrocery).Methods(http.MethodPut)
	r.HandleFunc("/groceries/{name}", grocery.DeleteGrocery).Methods(http.MethodDelete)

	log.Println("Listening ...")
	http.ListenAndServe(":8000", r)
}

package main

import (
	"log"
	"net/http"

	"github.com/gore/actions/grocery"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/allgroceries", grocery.AllGroceries).Methods("GET")
	r.HandleFunc("/groceries/{name}", grocery.SingleGrocery).Methods("GET")
	r.HandleFunc("/groceries", grocery.GroceriesToBuy).Methods("POST")
	r.HandleFunc("/groceries/{name}", grocery.UpdateGrocery).Methods("PUT")
	r.HandleFunc("/groceries/{name}", grocery.DeleteGrocery).Methods("DELETE")

	log.Println("Listening ...")
	http.ListenAndServe(":8000", r)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Grocery struct {
	Name     string `json: "name"`
	Quantity int    `json: "quantity"`
}

var groceries = []Grocery{
	{Name: "Almod Milk", Quantity: 2},
	{Name: "Apple", Quantity: 6},
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/allgroceries", AllGroceries).Methods("GET")
	r.HandleFunc("/groceries/{name}", SingleGrocery).Methods("GET")
	r.HandleFunc("/groceries", GroceriesToBuy).Methods("POST")
	r.HandleFunc("/groceries/{name}", UpdateGrocery).Methods("PUT")
	r.HandleFunc("/groceries/{name}", DeleteGrocery).Methods("DELETE")

	log.Println("Listening ...")
	http.ListenAndServe(":8000", r)
}

func AllGroceries(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: returnAllGroceries")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groceries)
}

func SingleGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	for _, grocery := range groceries {
		if grocery.Name == name {
			json.NewEncoder(w).Encode(grocery)
		}
	}
}

func GroceriesToBuy(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var grocery Grocery
	json.Unmarshal(reqBody, &grocery)
	groceries = append(groceries, grocery)

	json.NewEncoder(w).Encode(groceries)

}

func DeleteGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	for index, grocery := range groceries {
		if grocery.Name == name {
			groceries = append(groceries[:index], groceries[index+1:]...)
		}
	}

}

func UpdateGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	for index, grocery := range groceries {
		if grocery.Name == name {
			groceries = append(groceries[:index], groceries[index+1:]...)

			var updateGrocery Grocery

			json.NewDecoder(r.Body).Decode(&updateGrocery)
			groceries = append(groceries, updateGrocery)
			fmt.Println("Endpoint hit: UpdateGroceries")
			json.NewEncoder(w).Encode(updateGrocery)
			return
		}
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	gorillactx "github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	userID := gorillactx.Get(r, "user_id")
	json, err := db.GetOrders([]interface{}{userID.(string)})
	if err != nil {
		log.Printf("GetOrders DB error: %+v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, json)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json, err := db.GetProducts([]interface{}{})
	if err != nil {
		log.Printf("GetProducts DB error: %+v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, json)
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	userID := gorillactx.Get(r, "user_id")
	vars := mux.Vars(r)
	orderID := vars["id"]
	json, err := db.GetOrderProducts([]interface{}{orderID, userID})
	if err != nil {
		log.Printf("GetOrderByID DB error: %+v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, json)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	json, err := db.GetCategories([]interface{}{})
	if err != nil {
		log.Printf("GetCategories DB error: %+v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, json)
}

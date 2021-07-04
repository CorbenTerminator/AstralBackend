package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetHandler() http.Handler {
	r := mux.NewRouter()

	// r.HandleFunc("/welcome", Welcome).Methods("POST")

	//login and get jwt-token in cookie
	r.HandleFunc("/api/login", SignIn).Methods("POST")

	// r.HandleFunc("/api/refresh", RefreshToken).Methods("GET")

	r.Handle("/api/orders", AuthCheck(GetOrders)).Methods("GET")

	r.Handle("/api/orders/{id}", AuthCheck(GetOrderByID)).Methods("GET")

	r.Handle("/api/orders", AuthCheck(PostOrder)).Methods("POST")

	r.Handle("/api/products", AuthCheck(GetProducts)).Methods("GET")

	r.Handle("/api/products", AuthCheck(PostProduct)).Methods("POST")

	r.Handle("/api/prodorder/{id}", AuthCheck(PostOrderProducts)).Methods("POST")

	r.Handle("/api/categories", AuthCheck(GetCategories)).Methods("GET")

	return r
}

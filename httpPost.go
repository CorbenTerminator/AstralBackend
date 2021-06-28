package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	gorillactx "github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func PostOrder(w http.ResponseWriter, r *http.Request) {
	userID := gorillactx.Get(r, "user_id")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("PostOrder readBody error: %+v, Body: %s", err, string(body))
		http.Error(w, "Request body is not readable", http.StatusBadRequest)
		return
	}
	id, err := db.PostOrder(body, userID.(string))
	if err != nil {
		log.Printf("PostOrder DB error: %+v", err)
		http.Error(w, "Insert error", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, fmt.Sprintf("%d", id))
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("PostProduct readBody error: %+v, Body: %s", err, string(body))
		http.Error(w, "Request body is not readable", http.StatusBadRequest)
		return
	}
	id, err := db.PostProduct(body)
	if err != nil {
		log.Printf("PostProduct DB error: %+v", err)
		http.Error(w, "Insert error", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, fmt.Sprintf("%d", id))
}

func PostOrderProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("PostOrderProducts readBody error: %+v, Body: %s", err, string(body))
		http.Error(w, "Request body is not readable", http.StatusBadRequest)
		return
	}
	err = db.PostOrderProducts(body, orderID)
	if err != nil {
		log.Printf("PostOrderProducts DB error: %+v", err)
		http.Error(w, "Insert error", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Done")
}

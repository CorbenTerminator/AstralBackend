package main

import (
	"log"
	"net/http"

	mysql "AstralBackend/mysql"
)

const (
	dbConn = "user:user1234@tcp(mysql:3306)/shop?parseTime=true"
)

var db *mysql.Worker

func main() {
	var err error
	db, err = mysql.CreateWorker(dbConn)
	if err != nil {
		log.Printf("DB connection: %+v", err)
	}
	e := GetHandler()
	log.Fatal(http.ListenAndServe(":8080", e))
}

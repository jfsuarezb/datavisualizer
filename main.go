package main

import (
	
	"net/http"
	"log"
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var buyerURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"
var productURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products"
var transactionURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions"

func main() {
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	r.Get("/sync", SyncHandler)
	r.Get("/buyers", GetBuyers)

	err := http.ListenAndServe(":3000", r)

	if err != nil {

		log.Fatalln(err)

	} else {

		fmt.Println("App is running")

	}

}
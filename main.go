package main

import (
	
	"net/http"
	"log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var buyerURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"
var productURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products"
var transactionURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions"

func main() {
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	
	r.Get("/sync", syncHandler)

	http.ListenAndServe(":3000", r)

}

func syncHandler(w http.ResponseWriter, r *http.Request) {
	
	date := r.FormValue("date")

	buyersPayload, err := GetPayload(date, buyerURL)
		
	handleErr(w, err)

	buyersData := JSONHandler(buyersPayload)

	productsPayload, err := GetPayload(date, productURL)

	handleErr(w, err)

	productsData := CSVHandler(productsPayload)

	transactionsPayload, err := GetPayload(date, transactionURL)

	handleErr(w, err)

	transactionsData, err := NoStandHandler(transactionsPayload)

	handleErr(w, err)

	query := Concatenate(&buyersData, productsData, &transactionsData)

	query = "{\"set\":" + query + "}"

	resp, err := DGQuery(query)

	handleErr(w, err)

	w.Write([]byte(resp))

}

func handleErr(w http.ResponseWriter, err error) {

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("InternalServerError"))
		log.Fatalln(err)

	}

}
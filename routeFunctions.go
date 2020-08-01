package main

import (

	"net/http"
	"log"

)

var buyerURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"
var productURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products"
var transactionURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions"

func SyncHandler(w http.ResponseWriter, r *http.Request) {
	
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

	resp, err := DGPopulate(query)

	handleErr(w, err)

	w.Write([]byte(resp))

}

func GetBuyers(w http.ResponseWriter, r *http.Request) {

	resp, err := DGQueryBuyers()

	handleErr(w, err)

	w.Header().Set("Content-Type", "application/json")

	w.Write(resp)

}

func GetBuyer(w http.ResponseWriter, r *http.Request) {

	uid := r.FormValue("uid")

	resp, err := DGQueryBuyer(uid)

	handleErr(w, err)

	w.Header().Set("Content-Type", "application/json")

	w.Write(resp)

}

func handleErr(w http.ResponseWriter, err error) {

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("InternalServerError"))
		log.Fatalln(err)

	}

}
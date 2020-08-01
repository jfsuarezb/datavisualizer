package main

import (

	"net/http"
	"log"
	"os"

)

func SyncHandler(w http.ResponseWriter, r *http.Request) {
	
	buyerURL := os.Getenv("BUYERSLINK")
	productURL := os.Getenv("PRODUCTSLINK")
	transactionURL := os.Getenv("TRANSLINK")

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
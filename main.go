package main

import (
	"net/http"
	"fmt"
	"log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var buyerURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"
var productURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products"

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

	buyersBody, err := GetPayload(date, buyerURL)

	if err != nil {
		
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
		log.Fatalln(err)

	}

	w.Write([]byte("Succesful"))

	fmt.Println(buyersBody)

}
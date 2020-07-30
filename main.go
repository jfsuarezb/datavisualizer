package main

import (
	"net/http"
	"fmt"
	"log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

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

	buyersBody, err := GetBuyersBody(date)

	if err != nil {
		
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
		log.Fatalln(err)

	}

	w.Write([]byte("Succesful"))

	fmt.Println(buyersBody) 

}
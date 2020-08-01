package main

import (
	
	"net/http"
	"log"
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	r.Get("/sync", SyncHandler)
	r.Get("/buyers", GetBuyers)
	r.Get("/buyer", GetBuyer)

	err := http.ListenAndServe(":3000", r)

	if err != nil {

		log.Fatalln(err)

	} else {

		fmt.Println("App is running")

	}

}
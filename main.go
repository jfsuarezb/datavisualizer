package main

import (
	
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {
	
	godotenv.Load()
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", ServeHtml)
	r.Get("/index.css", ServeCss)
	r.Get("/index.js", ServeJs)
	
	r.Get("/sync", SyncHandler)
	r.Get("/buyers", GetBuyers)
	r.Get("/buyer", GetBuyer)

	http.ListenAndServe(":3000", r)

}
package main

import (
	"net/http"
	"fmt"
	"io/ioutil"

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

	url := fmt.Sprintf("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/​buyers?date=%s", date)

	fmt.Printf("\n\n%s", url)

	buyerResp, err := http.Get(url)		

	if err != nil {
		
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))

	}
	
	defer buyerResp.Body.Close()

	buyerBody, err := ioutil.ReadAll(buyerResp.Body)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))

	}

	fmt.Printf("%s", buyerBody)

	w.Write([]byte("Got the message"))

}
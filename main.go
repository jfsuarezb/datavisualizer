package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var buyerURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"

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

	buyersBody, err := getBuyersBody(date)

	if err != nil {
		
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
		log.Fatalln(err)

	}

	w.Write([]byte("Succesful"))

	fmt.Println(buyersBody)
}

func getBuyersBody(date string) (string, error) {
	
	request, err := http.NewRequest("GET", buyerURL, nil)
	request.Header.Set("date", date)

	if err != nil {
		return "", err
	}

	client := http.Client{}
	
	buyerResp, err := client.Do(request)

	if err != nil {
		return "", err
	}
	
	defer buyerResp.Body.Close()

	buyerBody, err := ioutil.ReadAll(buyerResp.Body)

	if err != nil {
		return "", err
	}

	return string(buyerBody), nil
}
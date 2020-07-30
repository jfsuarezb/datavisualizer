package main

import (
	
	"net/http"
	"io/ioutil"

)

var buyerURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"

func GetBuyersBody(date string) (string, error) {
	
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
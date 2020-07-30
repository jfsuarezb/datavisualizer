package main

import (
	
	"net/http"
	"io/ioutil"

)

func GetPayload(date string, url string) (string, error) {
	
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("date", date)

	if err != nil {
		return "", err
	}

	client := http.Client{}
	
	resp, err := client.Do(request)

	if err != nil {
		return "", err
	}
	
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil

}
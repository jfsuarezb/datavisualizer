package main

import (
	
	"net/http"
	"io/ioutil"

)

func GetPayload(date string, url string) ([]byte, error) {
	
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("date", date)

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}
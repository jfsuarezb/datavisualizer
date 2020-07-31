package main

import (

	"encoding/json"
	"strings"

)

func JSONHandler(payload []byte) ([]map[string]string) {

	var data []map[string]string

	json.Unmarshal(payload, &data)

	for _, item := range data {

		item["bid"] = item["id"]
		delete(item, "id")

	}

	return data

}

func CSVHandler(payload []byte) ([]map[string]string) {

	s := strings.Split(string(payload), "\n")

	var arr []string

	var data []map[string]string

	placeholder := make(map[string]string)

	for _, ch := range s {

		arr = strings.Split(ch, "'")

		if len(arr) == 3 {

			placeholder["pid"] = arr[0]
			placeholder["name"] = arr[1]
			placeholder["price"] = arr[2]

		}

		data = append(data, placeholder)

	}

	return data

}
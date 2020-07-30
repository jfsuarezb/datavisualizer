package main

import (

	"encoding/json"

)

func JSONHandler(payload string) ([]map[string]string) {

	var data []map[string]string

	json.Unmarshal([]byte(payload), &data)

	for _, item := range data {

		item["bid"] = item["id"]
		delete(item, "id")

	}

	return data

}
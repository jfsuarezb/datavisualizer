package main

import (

	"encoding/json"
	"strings"
	"regexp"

)

func JSONHandler(payload []byte) ([]map[string]interface{}) {

	var data []map[string]interface{}

	json.Unmarshal(payload, &data)

	for _, item := range data {

		item["bid"] = item["id"]
		delete(item, "id")

	}

	return data

}

func CSVHandler(payload []byte) ([]map[string]interface{}) {

	s := strings.Split(string(payload), "\n")

	var arr []string

	var data []map[string]interface{}

	for _, ch := range s {

		placeholder := make(map[string]interface{})

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

func NoStandHandler(payload []byte) ([]map[string]interface{}, error) {

	s := strings.Split(string(payload), "#")

	var data []map[string]interface{}

	bidr, err := regexp.Compile("[0-9a-f]{6,8}")

	if err != nil {

		return nil, err

	}

	r, err := regexp.Compile("([\\d\\.])+")

	if err != nil {

		return nil, err

	}

	dotr, err := regexp.Compile("\\.")

	if err != nil {

		return nil, err

	}

	devr, err := regexp.Compile("ios|android|mac|linux|windows")

	if err != nil {

		return nil, err

	}

	parr, err := regexp.Compile("\\(([^)]+)\\)")

	if err != nil {

		return nil, err

	}

	pidr, err := regexp.Compile("(\\w+)")

	if err != nil {

		return nil, err

	}


	for _, ch := range s {

		placeholder := make(map[string]interface{})

		placeholder["tid"] = string([]rune(ch)[0:12])

		if bidr.MatchString(ch) {

			placeholder["bid"] = bidr.FindAllString(ch, -1)[1]

		}

		match := r.FindAllString(ch, -1)

		for _, posip := range match {
			
			if dotr.MatchString(posip) {

				placeholder["ip"] = posip

				break

			}
		}

		placeholder["dev"] = devr.FindString(ch)

		pids := parr.FindString(ch)

		placeholder["pids"] = pidr.FindAllString(pids, -1)

		data = append(data, placeholder)

	}

	return data, nil

}
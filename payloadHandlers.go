package main

import (

	"encoding/json"
	"strings"
	"regexp"

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

	for _, ch := range s {

		placeholder := make(map[string]string)

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

func NoStandHandler(payload []byte) ([]map[string]interface{}) {

	s := strings.Split(string(payload), "#")

	var data []map[string]interface{}

	r, _ := regexp.Compile("([\\d\\.])+")
	dotr, _ := regexp.Compile("\\.")
	devr, _ := regexp.Compile("ios|android|mac|linux|windows")
	parr, _ := regexp.Compile("\\(([^)]+)\\)")
	pidr, _ := regexp.Compile("(\\w+)")


	for _, ch := range s {

		placeholder := make(map[string]interface{})

		placeholder["tid"] = string([]rune(ch)[0:11])
		placeholder["bid"] = string([]rune(ch)[12:19])

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

	return data

}
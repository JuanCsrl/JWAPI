package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Month struct {
	Agosto []struct {
		ID        string `json:"id"`
		Data      string `json:"data"`
		Versiculo string `json:"versiculo"`
		Text      string `json:"text"`
	} `json:"agosto"`
}

func getFile() Month {
	jsonFile, err := os.Open("month.json")
	if err != nil {
		logsError(err)
	}
	logsInfo("Successfully Opened month.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data Month
	json.Unmarshal(byteValue, &data)
	return data
}

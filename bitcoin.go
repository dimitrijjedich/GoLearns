package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiData struct {
	Time struct {
		Updated    string    `json:"updated"`
		UpdatedISO time.Time `json:"updatedISO"`
		Updateduk  string    `json:"updateduk"`
	} `json:"time"`
	Disclaimer string `json:"disclaimer"`
	ChartName  string `json:"chartName"`
	Bpi        struct {
		USD struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"USD"`
		GBP struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"GBP"`
		EUR struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"EUR"`
	} `json:"bpi"`
}

func main() {
	var apiUrl string = "https://api.coindesk.com/v1/bpi/currentprice.json"
	var client *http.Client = &http.Client{}
	var result, err = client.Get(apiUrl)
	if err != nil {
		fmt.Printf("An error occured:\n")
		fmt.Print(err)
		return
	}
	var data []byte
	data, err = io.ReadAll(result.Body)
	if err != nil {
		fmt.Printf("Something went wrong during parsing:\n")
		fmt.Print(err)
		return
	}
	var bitcoinData ApiData
	json.Unmarshal(data, &bitcoinData)
	fmt.Printf("Data from: %v \n", bitcoinData.Time.Updated)
	fmt.Printf("Data from: %v", bitcoinData.Bpi.EUR.Rate)
}

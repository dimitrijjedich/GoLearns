package main

import (
	"fmt"
	"net/http"
)

func main() {
	var apiUrl string = "https://api.coindesk.com/v1/bpi/currentprice.json"
	var client *http.Client = &http.Client{}
	var result, err = client.Get(apiUrl)
	if err != nil {
		fmt.Printf("An error occured:\n")
		fmt.Prinf(err)
		return
	}
}

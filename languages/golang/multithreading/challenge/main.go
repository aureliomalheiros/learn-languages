package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type AddressResult struct {
	Address      string  `json:"address"`
	Neighborhood string  `json:"neighborhood"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	API          string  `json:"api"`
	ResponseTime float64 `json:"response_time"`
}

func requestAddressData(apiURL, cep string, resultChan chan<- AddressResult, wg *sync.WaitGroup) {
	defer wg.Done()

	startTime := time.Now()

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Get(fmt.Sprintf(apiURL, cep))
	if err != nil {
		fmt.Printf("Error making request to %s: %v\n", apiURL, err)
		return
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body from %s: %v\n", apiURL, err)
		return
	}

	var result AddressResult
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		fmt.Printf("Error decoding JSON from %s: %v\n", apiURL, err)
		return
	}

	result.API = apiURL
	result.ResponseTime = time.Since(startTime).Seconds()

	resultChan <- result
}

func main() {
	var wg sync.WaitGroup
	resultChan := make(chan AddressResult, 2)

	cep := "01153000"

	apiURLs := []string{
		"https://brasilapi.com.br/api/cep/v1/%s",
		"http://viacep.com.br/ws/%s/json/",
	}

	for _, apiURL := range apiURLs {
		wg.Add(1)
		go requestAddressData(apiURL, cep, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	var fastestResult AddressResult
	fastestResult.ResponseTime = 1.0 // Max response time

	for result := range resultChan {
		if result.ResponseTime < fastestResult.ResponseTime {
			fastestResult = result
		}
	}

	if fastestResult.ResponseTime == 1.0 {
		fmt.Println("No response received within the timeout.")
	} else {
		fmt.Printf("Fastest result from API %s:\n", fastestResult.API)
		fmt.Printf("Address Data: %+v\n", fastestResult)
	}
}

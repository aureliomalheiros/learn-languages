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

type ViaCEPResult struct {
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func requestAddressData(apiURL, cep string, resultChan chan<- AddressResult, wg *sync.WaitGroup) {
	defer wg.Done()

	startTime := time.Now()

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	reqURL := fmt.Sprintf(apiURL, cep)

	resp, err := client.Get(reqURL)
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
	switch apiURL {
	case "https://brasilapi.com.br/api/cep/v1/%s":
		err = json.Unmarshal(responseData, &result)
	case "http://viacep.com.br/ws/%s/json/":
		var viaCEPResult ViaCEPResult
		err = json.Unmarshal(responseData, &viaCEPResult)
		result = AddressResult{
			Address:      viaCEPResult.Logradouro,
			Neighborhood: viaCEPResult.Bairro,
			City:         viaCEPResult.Localidade,
			State:        viaCEPResult.Uf,
			API:          reqURL, // Define the API URL here
			ResponseTime: time.Since(startTime).Seconds(),
		}
	}

	if err != nil {
		fmt.Printf("Error decoding JSON from %s: %v\n", apiURL, err)
		return
	}

	// Print response time for debugging
	fmt.Printf("Response time for API %s: %f seconds\n", apiURL, time.Since(startTime).Seconds())

	resultChan <- result
}

func main() {
	cep := "01153000"
	apiURLs := []string{
		"https://brasilapi.com.br/api/cep/v1/%s",
		"http://viacep.com.br/ws/%s/json/",
	}

	resultChan := make(chan AddressResult, len(apiURLs))
	var wg sync.WaitGroup

	for _, apiURL := range apiURLs {
		wg.Add(1)
		go requestAddressData(apiURL, cep, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	var fastestResult AddressResult
	fastestTime := 1.0

	for result := range resultChan {
		if result.ResponseTime < fastestTime {
			fastestResult = result
			fastestTime = result.ResponseTime
		}
	}

	if fastestTime == 1.0 {
		fmt.Println("All requests timed out.")
	} else {
		fmt.Printf("Fastest result from API %s:\n", fastestResult.API)
		fmt.Printf("Address Data: %+v\n", fastestResult)
	}
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_"os"
	"time"
)

const (
	serverURL      = "http://localhost:8080/cotacao"
	requestTimeout = 300 * time.Millisecond
)

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", serverURL, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to get valid response: %v", resp.Status)
	}

	var cotacao CotacaoResponse
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	err = ioutil.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %s", cotacao.Bid)), 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	log.Println("Cotação salva com sucesso em cotacao.txt")
}

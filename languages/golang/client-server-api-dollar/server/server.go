package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	apiURL        = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	queryTimeout  = 500 * time.Millisecond // Aumentando o timeout para 500ms
	insertTimeout = 10 * time.Millisecond
)

type Cotacao struct {
	Bid string `json:"bid"`
}

type ApiResponse struct {
	USDBRL Cotacao `json:"USDBRL"`
}

func getCotacao(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return "", err
	}

	return apiResponse.USDBRL.Bid, nil
}

func saveCotacao(ctx context.Context, db *sql.DB, bid string) error {
	query := "INSERT INTO cotacoes (bid, timestamp) VALUES (?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, bid, time.Now().Unix())
	return err
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), queryTimeout)
	defer cancel()

	bid, err := getCotacao(ctx)
	if err != nil {
		http.Error(w, "Failed to get cotacao", http.StatusInternalServerError)
		log.Printf("Error getting cotacao: %v", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), insertTimeout)
	defer cancel()

	db, err := sql.Open("sqlite3", os.Getenv("DB_NAME"))
	if err != nil {
		http.Error(w, "Failed to open database", http.StatusInternalServerError)
		log.Printf("Error opening database: %v", err)
		return
	}
	defer db.Close()

	if err := saveCotacao(ctx, db, bid); err != nil {
		http.Error(w, "Failed to save cotacao", http.StatusInternalServerError)
		log.Printf("Error saving cotacao: %v", err)
		return
	}

	response := map[string]string{"bid": bid}
	json.NewEncoder(w).Encode(response)
}

func setupDatabase() error {
	db, err := sql.Open("sqlite3", os.Getenv("DB_NAME"))
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT,
		timestamp INTEGER
	)`
	_, err = db.Exec(query)
	return err
}

func main() {
	if err := setupDatabase(); err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	http.HandleFunc("/cotacao", cotacaoHandler)
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

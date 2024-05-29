package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        os.Exit(1)
    }

    // Verificar se a chave da API foi carregada
    apiKey := os.Getenv("WEATHER_API_KEY")
    if apiKey == "" {
        fmt.Println("API key not found")
        os.Exit(1)
    }
    weatherAPIKey = apiKey

    os.Exit(m.Run())
}

func TestGetWeatherValidCEP(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/weather/01001000", nil) // CEP válido
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetWeatherInvalidCEP(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/weather/1234567", nil) // CEP inválido (menos de 8 dígitos)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestGetWeatherNotFoundCEP(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/weather/99999999", nil) // CEP que provavelmente não existe
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)
}

func setupRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.GET("/weather/:cep", getWeather)
    return router
}

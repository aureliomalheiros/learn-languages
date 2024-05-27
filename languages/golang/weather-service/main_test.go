package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

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

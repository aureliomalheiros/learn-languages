package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	viaCepURL     = "https://viacep.com.br/ws/%s/json/"
	weatherAPIURL = "http://api.weatherapi.com/v1/current.json?key=%s&q=%s"
)

var weatherAPIKey = os.Getenv("WEATHER_API_KEY")

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func main() {
	r := gin.Default()
	r.GET("/weather/:cep", getWeather)
	r.Run()
}

func getWeather(c *gin.Context) {
	cep := c.Param("cep")

	if len(cep) != 8 {
		log.Printf("Invalid CEP length: %s\n", cep)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
		return
	}

	viaCepResponse, err := getCityFromCep(cep)
	if err != nil {
		log.Printf("Error getting city from CEP: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "can not find zipcode"})
		return
	}

	log.Printf("Retrieved city: %s\n", viaCepResponse.Localidade)

	weatherResponse, err := getWeatherFromCity(viaCepResponse.Localidade)
	if err != nil {
		log.Printf("Error getting weather from city: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error retrieving weather"})
		return
	}

	c.JSON(http.StatusOK, weatherResponse)
}

func getCityFromCep(cep string) (ViaCepResponse, error) {
	var viaCepResponse ViaCepResponse
	resp, err := http.Get(fmt.Sprintf(viaCepURL, cep))
	if err != nil {
		log.Printf("Error making request to ViaCep: %v\n", err)
		return viaCepResponse, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("ViaCep API returned non-OK status: %d\n", resp.StatusCode)
		return viaCepResponse, fmt.Errorf("failed to get city from cep")
	}

	err = json.NewDecoder(resp.Body).Decode(&viaCepResponse)
	if err != nil {
		log.Printf("Error decoding ViaCep response: %v\n", err)
		return viaCepResponse, err
	}

	if viaCepResponse.Localidade == "" {
		log.Printf("City not found for CEP: %s\n", cep)
		return viaCepResponse, fmt.Errorf("city not found for cep")
	}

	return viaCepResponse, nil
}

func getWeatherFromCity(city string) (WeatherResponse, error) {
	var weatherAPIResponse WeatherAPIResponse
	encodedCity := url.QueryEscape(city) // Encode the city to handle special characters
	url := fmt.Sprintf(weatherAPIURL, weatherAPIKey, encodedCity)
	log.Printf("Requesting weather data from URL: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making request to WeatherAPI: %v\n", err)
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("WeatherAPI returned non-OK status: %d\n", resp.StatusCode)
		return WeatherResponse{}, fmt.Errorf("failed to get weather for city")
	}

	err = json.NewDecoder(resp.Body).Decode(&weatherAPIResponse)
	if err != nil {
		log.Printf("Error decoding WeatherAPI response: %v\n", err)
		return WeatherResponse{}, err
	}

	tempC := weatherAPIResponse.Current.TempC
	tempF := tempC*1.8 + 32
	tempK := tempC + 273.15

	return WeatherResponse{TempC: tempC, TempF: tempF, TempK: tempK}, nil
}

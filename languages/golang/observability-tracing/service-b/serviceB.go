package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "github.com/gorilla/mux"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/zipkin"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    "context"
    "log"
)

type ZipCodeRequest struct {
    CEP string `json:"cep"`
}

type ZipCodeResponse struct {
    City   string  `json:"city"`
    TempC  float64 `json:"temp_C"`
    TempF  float64 `json:"temp_F"`
    TempK  float64 `json:"temp_K"`
}

type ViaCEPResponse struct {
    Localidade string `json:"localidade"`
}

type WeatherAPIResponse struct {
    Current struct {
        TempC float64 `json:"temp_c"`
    } `json:"current"`
}

const weatherAPIKey = "8665baf001c34b0c921185735242705" 

func main() {
    // Initialize Zipkin exporter
    exporter, err := zipkin.New("http://localhost:9411/api/v2/spans")
    if err != nil {
        log.Fatalf("failed to initialize zipkin exporter %v", err)
    }

    // Create trace provider
    tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
    defer func() { _ = tp.Shutdown(context.Background()) }()
    otel.SetTracerProvider(tp)

    r := mux.NewRouter()
    r.HandleFunc("/cep", handleZipCodeRequest).Methods("POST")

    log.Println("Service B is running on port 8081")
    http.ListenAndServe(":8081", r)
}

func handleZipCodeRequest(w http.ResponseWriter, r *http.Request) {
    var zipCodeRequest ZipCodeRequest
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &zipCodeRequest)

    log.Printf("Received CEP: %s", zipCodeRequest.CEP)

    // Validate CEP
    if !isValidCEP(zipCodeRequest.CEP) {
        log.Println("Invalid CEP format")
        http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
        return
    }

    tracer := otel.Tracer("service-b")
    ctx, span := tracer.Start(context.Background(), "handleZipCodeRequest")
    defer span.End()

    // Find location by CEP
    location, err := findLocationByCEP(ctx, zipCodeRequest.CEP)
    if err != nil {
        log.Printf("Error finding location: %v", err)
        http.Error(w, "can not find zipcode", http.StatusNotFound)
        return
    }

    // Get temperature
    tempC, err := getTemperatureByLocation(ctx, location)
    if err != nil {
        log.Printf("Error getting temperature: %v", err)
        http.Error(w, "can not find temperature", http.StatusInternalServerError)
        return
    }

    tempF := tempC*1.8 + 32
    tempK := tempC + 273.15

    response := ZipCodeResponse{
        City:  location,
        TempC: tempC,
        TempF: tempF,
        TempK: tempK,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func isValidCEP(cep string) bool {
    re := regexp.MustCompile(`^\d{8}$`)
    return re.MatchString(cep)
}

func findLocationByCEP(ctx context.Context, cep string) (string, error) {
    tracer := otel.Tracer("service-b")
    _, span := tracer.Start(ctx, "findLocationByCEP")
    defer span.End()

    resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var viaCEPResponse ViaCEPResponse
    body, _ := ioutil.ReadAll(resp.Body)
    json.Unmarshal(body, &viaCEPResponse)

    if viaCEPResponse.Localidade == "" {
        return "", fmt.Errorf("location not found")
    }

    return viaCEPResponse.Localidade, nil
}

func getTemperatureByLocation(ctx context.Context, location string) (float64, error) {
    tracer := otel.Tracer("service-b")
    _, span := tracer.Start(ctx, "getTemperatureByLocation")
    defer span.End()

    resp, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", weatherAPIKey, location))
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    var weatherAPIResponse WeatherAPIResponse
    body, _ := ioutil.ReadAll(resp.Body)
    json.Unmarshal(body, &weatherAPIResponse)

    return weatherAPIResponse.Current.TempC, nil
}


package main

import (
    "bytes"
    "encoding/json"
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

    log.Println("Service A is running on port 8080")
    http.ListenAndServe(":8080", r)
}

func handleZipCodeRequest(w http.ResponseWriter, r *http.Request) {
    var zipCodeRequest ZipCodeRequest
    body, _ := ioutil.ReadAll(r.Body)
    r.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // Reabastece o corpo da requisição

    json.Unmarshal(body, &zipCodeRequest)

    log.Printf("Received CEP: %s", zipCodeRequest.CEP)

    // Validate CEP
    if !isValidCEP(zipCodeRequest.CEP) {
        log.Println("Invalid CEP format")
        http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
        return
    }

    tracer := otel.Tracer("service-a")
    _, span := tracer.Start(context.Background(), "handleZipCodeRequest")
    defer span.End()

    // Call Service B
    log.Println("Calling Service B")
    resp, err := http.Post("http://service-b:8081/cep", "application/json", bytes.NewBuffer(body))
    if err != nil {
        log.Printf("Error calling Service B: %v", err)
        http.Error(w, "can not find zipcode", http.StatusNotFound)
        return
    }
    defer resp.Body.Close()

    log.Printf("Service B response status: %d", resp.StatusCode)
    if resp.StatusCode != http.StatusOK {
        body, _ = ioutil.ReadAll(resp.Body)
        log.Printf("Service B response body: %s", string(body))
        http.Error(w, "can not find zipcode", resp.StatusCode)
        return
    }

    var zipCodeResponse ZipCodeResponse
    body, _ = ioutil.ReadAll(resp.Body)
    json.Unmarshal(body, &zipCodeResponse)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(zipCodeResponse)
}

func isValidCEP(cep string) bool {
    re := regexp.MustCompile(`^\d{8}$`)
    return re.MatchString(cep)
}

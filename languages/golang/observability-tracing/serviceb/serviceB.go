package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/semconv/v1.7.0"
)

type WeatherResponse struct {
    City  string  `json:"city"`
    TempC float64 `json:"temp_C"`
    TempF float64 `json:"temp_F"`
    TempK float64 `json:"temp_K"`
}

func main() {
    ctx := context.Background()

    exporter, err := otlptracehttp.New(ctx)
    if err != nil {
        log.Fatalf("failed to create OTLP trace exporter: %v", err)
    }

    tp := trace.NewTracerProvider(
        trace.WithBatcher(exporter),
        trace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String("service-b"),
        )),
    )
    otel.SetTracerProvider(tp)
    defer func() {
        if err := tp.Shutdown(context.Background()); err != nil {
            log.Fatalf("failed to shutdown tracer provider: %v", err)
        }
    }()

    http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
        tracer := tp.Tracer("service-b-tracer")
        ctx, span := tracer.Start(context.Background(), "weather-handler")
        defer span.End()

        cep := r.URL.Query().Get("cep")
        if len(cep) != 8 {
            w.WriteHeader(http.StatusUnprocessableEntity)
            json.NewEncoder(w).Encode(map[string]string{"message": "invalid zipcode"})
            return
        }

        // Fake data for the example
        city := "São Paulo"
        tempC := 28.5
        tempF := tempC*1.8 + 32
        tempK := tempC + 273.15

        response := WeatherResponse{
            City:  city,
            TempC: tempC,
            TempF: tempF,
            TempK: tempK,
        }

        span.SetAttributes(
            attribute.String("city", city),
            attribute.Float64("tempC", tempC),
        )

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(response)
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
}

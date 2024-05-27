# Observability and trancing

# Temperature by Zip Code Service

This project consists of two services (Service A and Service B) that work together to receive a zip code, identify the corresponding city, and return the current temperature of that city in Celsius, Fahrenheit, and Kelvin. The system implements distributed tracing using OpenTelemetry (OTEL) and Zipkin to monitor and trace requests.

## Prerequisites

- Docker and Docker Compose installed on your machine.
- An API key for WeatherAPI (you can register and obtain a key at https://www.weatherapi.com/).

## Project Structure

```markdown
observability-tracing
├── service-a
│   ├── Dockerfile
│   └── main.go
├── service-b
│   ├── Dockerfile
│   └── main.go
├── docker-compose.yml
└── run.sh
```

## Configure the Env

const weatherAPIKey = "YOUR_WEATHER_API_KEY"

### Running the project

```bash
chmod +x run.sh
```

```bash
sudo ./run.sh
```

Requests:


```bash
curl -X POST -H "Content-Type: application/json" -d '{"cep": "29902555"}' http://localhost:8080/cep\ 
```

```bash
#Result
{"city":"Linhares","temp_C":22.7,"temp_F":72.86,"temp_K":295.84999999999997}
```

Access Zipkin:

```bash
http://localhost:9411
```


#!/bin/bash

set -e

echo "Building and starting the services using Docker Compose..."

docker-compose up --build

echo "Waiting for the services to start..."

sleep 10

echo "Services are up and running."

echo "You can access Service A at http://localhost:8080/cep"
echo "You can access Zipkin at http://localhost:9411"

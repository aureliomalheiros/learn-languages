# Weather service

```markdown
weather-service/
├── main.go
├── test/
│   └── main_test.go
├── Dockerfile
├── docker-comopose.yaml
├── .env
├── go.mod
├── go.sum
└── README.md
```

## Running the project

Local:

```bash
go run main.go
./main
```

```bash
curl http://localhost:8080/weather/YOUR_CEP
```

Docker Compose:

> [!Note]
> It's necessary modify value in file `.env` with your token.

```bash
sudo docker-compose up --build
```

Example:

```bash
curl http://localhost:8080/weather/65055540

{"temp_C":28,"temp_F":82.4,"temp_K":301.15}
```

Cloud Run:

[My URL in Cloud Run](https://weather-service-y4g7b3akmq-ue.a.run.app/weather/65055540)


```json
{"temp_C":28,"temp_F":82.4,"temp_K":301.15}
```


### Tests

You can running tests in GH Actions or in local machine.

```bash
go test -v 
```
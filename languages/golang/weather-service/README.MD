# Weather service

```markdown
weather-service/
├── main.go
├── test/
│   └── main_test.go
├── Dockerfile
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


Example:

```bash
curl http://localhost:8080/weather/65055540

{"temp_C":28,"temp_F":82.4,"temp_K":301.15}
```

Cloud Run:

```bash
https://weather-service-y4g7b3akmq-ue.a.run.app/weather/65055540
```

```json
{"temp_C":28,"temp_F":82.4,"temp_K":301.15}
```


### Tests
# Client Server dollar quote

:money_with_wings: This repository is a project for requesting the dollar quote.

![Status](https://img.shields.io/badge/status-progress-blue)

## Requirements

- Golang
- Containers (Podmand or Docker)
- Database SQLite

### Explain the Project - Challenge

Client:

- Will make requests to receive the USD exchange rate
- Maximum timeout of 300ms to receive the exchange rate
- Should save the exchange rates in a file named "cotacao.txt" (Create a folder)
- Will receive only the JSON bind

Server:

- The server will make requests to the API
- You will receive a JSON response
- Will record the USD exchange rates in an SQLite database
- The maximum timeout to call the API is 200ms
- Maximum timeout to persist data in the database is 10ms
- Endpoint: `/cotacao`
- Port: `8080`

All three contexts should log an error if the execution time is insufficient.

### Running

```bash
client-server-api-dollar
├── client/
│   └── client.go
├── server/
│   └── server.go
│   └── Dockerfile 
└── docker-compose.yaml
└── README.MD
```

Execute the command in root folder:

```bash
sudo docker-compose up --build
```

The bellow command exec the client and You will generate file `cotacao.txt`

```bash
cd client
go run client.go
```

If you will want generate quote dollars in web client it's possible use URL: `http://localhost:8080/cotacao`


package main

import (
    "flag"
    "fmt"
    "net/http"
    "sync"
    "time"
)

func main() {
    // Definindo flags de linha de comando
    url := flag.String("url", "", "URL do serviço a ser testado")
    requests := flag.Int("requests", 0, "Número total de requests")
    concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")
    flag.Parse()

    // Verificar se os parâmetros obrigatórios foram fornecidos
    if *url == "" || *requests == 0 {
        flag.PrintDefaults()
        return
    }

    // Iniciar o teste de carga
    start := time.Now()
    fmt.Println("Iniciando teste de carga...")

    var wg sync.WaitGroup
    result := make(chan int, *requests)

    for i := 0; i < *concurrency; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < *requests / *concurrency; j++ {
                makeRequest(*url, result)
            }
        }()
    }

    go func() {
        wg.Wait()
        close(result)
    }()

    // Processar resultados
    total := 0
    success := 0
    statusCodes := make(map[int]int)

    for res := range result {
        total++
        if res == 200 {
            success++
        }
        statusCodes[res]++
    }

    // Gerar relatório
    elapsed := time.Since(start)
    fmt.Println("Teste de carga concluído.")
    fmt.Printf("Tempo total gasto: %v\n", elapsed)
    fmt.Printf("Total de requests: %d\n", total)
    fmt.Printf("Requests com status 200: %d\n", success)
    fmt.Println("Distribuição de códigos de status:")
    for code, count := range statusCodes {
        fmt.Printf("Status %d: %d\n", code, count)
    }
}

func makeRequest(url string, result chan<- int) {
    response, err := http.Get(url)
    if err != nil {
        result <- 500 // Error 500 - Internal Server Error
        return
    }
    defer response.Body.Close()
    result <- response.StatusCode
}

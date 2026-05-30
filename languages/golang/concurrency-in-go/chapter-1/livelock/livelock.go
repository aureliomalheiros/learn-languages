package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type Transaction struct {
    id     int
    value  int
    retry  int
}

func process(t *Transaction, mu *sync.Mutex, data map[int]int) {
    for t.retry < 100 {
        mu.Lock()
        
        // Simula condição que sempre falha
        if rand.Intn(2) == 0 {
            mu.Unlock()
            t.retry++
            
            // LIVELOCK: backoff aleatório mas nunca resolve
            backoff := time.Duration(rand.Intn(10)) * time.Millisecond
            time.Sleep(backoff)
            continue
        }
        
        // Processa
        data[t.id] = t.value
        mu.Unlock()
        fmt.Printf("Transação %d concluída\n", t.id)
        return
    }
    fmt.Printf("Transação %d falhou após %d tentativas\n", t.id, t.retry)
}

func main() {
    var mu sync.Mutex
    data := make(map[int]int)
    
    // Múltiplas transações competindo pela mesma condição
    for i := 0; i < 5; i++ {
        go process(&Transaction{id: i, value: i * 100}, &mu, data)
    }
    
    time.Sleep(5 * time.Second)
}

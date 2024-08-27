package main

import (
    "fmt"
    "strconv"
    "sync"
    "time"
)

func worker(start, end int, password string, wg *sync.WaitGroup, result *string) {
    defer wg.Done()
    for i := start; i <= end; i++ {
        attempt := fmt.Sprintf("%08d", i)
        if attempt == password {
            *result = attempt
            return
        }
    }
}

func crackPasswordParallel(password string, numThreads int) string {
    var result string
    var wg sync.WaitGroup
    start := time.Now()
    chunkSize := 100000000 / numThreads

    for i := 0; i < numThreads; i++ {
        wg.Add(1)
        go worker(i*chunkSize, (i+1)*chunkSize-1, password, &wg, &result)
    }

    wg.Wait()
    fmt.Printf("Password found: %s\n", result)
    fmt.Printf("Time taken: %v\n", time.Since(start))
    return result
}

func main() {
    password := "12345678" // A senha a ser quebrada
    numThreads := 4        // Número de threads a serem usadas
    crackPasswordParallel(password, numThreads)
}

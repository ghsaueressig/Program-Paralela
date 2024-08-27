package main

import (
    "fmt"
    "strconv"
    "time"
)

func crackPasswordSequential(password string) string {
    start := time.Now()
    for i := 0; i <= 99999999; i++ {
        attempt := fmt.Sprintf("%08d", i)
        if attempt == password {
            fmt.Printf("Password found: %s\n", attempt)
            fmt.Printf("Time taken: %v\n", time.Since(start))
            return attempt
        }
    }
    return ""
}

func main() {
    password := "12345678" // A senha a ser quebrada
    crackPasswordSequential(password)
}

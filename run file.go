package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
    password := r.URL.Query().Get("password")
    numThreads, _ := strconv.Atoi(r.URL.Query().Get("threads"))
    
    if password == "" || numThreads <= 0 {
        fmt.Fprintf(w, "Please provide a valid password and number of threads.")
        return
    }

    fmt.Fprintf(w, "Cracking password: %s with %d threads\n", password, numThreads)
    crackPasswordParallel(password, numThreads)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}

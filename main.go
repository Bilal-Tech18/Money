package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <url>")
        os.Exit(1)
    }
    siteURL := os.Args[1]

    stats := startCrawling(siteURL)
    if stats != nil {
        fmt.Printf("Total: %d pages, %d inputs, %d inputs hidden\n",
            stats.TotalPages, stats.TotalInputs, stats.TotalHiddenInputs)
    }
}

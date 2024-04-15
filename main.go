package main

import (
    "fmt"
    "os"
)

func main() {
    // Vérifier si un argument URL a été fourni
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <url>")
        os.Exit(1)
    }
    siteURL := os.Args[1] // L'URL sera le premier argument de la ligne de commande

	pageCount := startCrawling(siteURL)
    fmt.Printf("Total pages visited: %d\n", pageCount)
}


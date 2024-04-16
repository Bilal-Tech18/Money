package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	excelFlag := flag.Bool("excel", false, "Enable to write output to an Excel file")
	flag.Parse() // Cela doit être appelé après avoir défini tous les flags et avant de lire les flags

	// Après flag.Parse(), os.Args ne contiendra que les arguments non flag
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go [--excel] <url>")
		os.Exit(1)
	}
	siteURL := flag.Arg(0)

	fmt.Printf("URL to crawl: %s\n", siteURL)
	fmt.Printf("Write to Excel flag is set to: %v\n", *excelFlag) // Pour vérifier l'état du flag

	stats := startCrawling(siteURL, *excelFlag)
	if stats != nil {
		fmt.Printf("Total: %d pages, %d inputs, %d inputs hidden\n",
			stats.TotalPages, stats.TotalInputs, stats.TotalHiddenInputs)
	}
}

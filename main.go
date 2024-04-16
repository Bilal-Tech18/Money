package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	excelFlag := flag.Bool("excel", false, "Enable to write output to an Excel file")
	flag.Parse()

	// Vérifiez si l'URL à crawler est fournie en argument
	if flag.NArg() < 1 {
		fmt.Println("Usage: go run main.go [--excel] <url>")
		os.Exit(1)
	}
	siteURL := flag.Arg(0)

	fmt.Printf("URL to crawl: %s\n", siteURL)
	fmt.Printf("Write to Excel flag is set to: %v\n", *excelFlag)

	stats := startCrawling(siteURL, *excelFlag)
	if stats != nil {
		fmt.Printf("Total: %d pages, %d inputs, %d inputs hidden\n",
			stats.TotalPages, stats.TotalInputs, stats.TotalHiddenInputs)
	}
}

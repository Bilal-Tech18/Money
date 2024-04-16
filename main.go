package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	helpFlag := flag.Bool("help", false, "Display usage information")
	excelFlag := flag.Bool("excel", false, "Enable to write output to an Excel file")
	txtFlag := flag.Bool("txt", false, "Enable to write output to a text file")

	flag.Parse()

	if *helpFlag {
		printUsage()
		os.Exit(0)
	}

	// Vérifiez si l'URL à crawler est fournie en argument
	if flag.NArg() < 1 {
		fmt.Println("Usage: go run main.go [--excel|--txt] <url>")
		os.Exit(1)
	}
	siteURL := flag.Arg(0)

	fmt.Printf("URL to crawl: %s\n", siteURL)

	if *excelFlag && *txtFlag {
		fmt.Println("Error: Please choose only one output option, either --excel or --txt")
		os.Exit(1)
	}

	var stats *Stats
	if *excelFlag {
		stats = startCrawling(siteURL, true)
	} else if *txtFlag {
		stats = startCrawling(siteURL, false)
	} else {
		fmt.Println("Error: Please specify either --excel or --txt")
		os.Exit(1)
	}

	if stats != nil {
		fmt.Printf("Total: %d pages, %d inputs, %d inputs hidden\n",
			stats.TotalPages, stats.TotalInputs, stats.TotalHiddenInputs)
	}
}

// Fonction pour afficher le message d'aide
func printUsage() {
	fmt.Println("Usage: go run main.go [--excel|--txt|--help] <url>")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

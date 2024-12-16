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
	jsonFlag := flag.Bool("json", false, "Enable to write output to a JSON file")
	verboseFlag := flag.Bool("verbose", false, "Print each visited URL to the terminal")

	flag.Parse()

	if *helpFlag {
		printUsage()
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		printUsage()
		os.Exit(1)
	}
	siteURL := flag.Arg(0)

	fmt.Printf("URL to crawl: %s\n", siteURL)

	// Vérifier qu'une seule option de sortie est sélectionnée
	if (*excelFlag && *txtFlag) || (*excelFlag && *jsonFlag) || (*txtFlag && *jsonFlag) {
		fmt.Println("Error: Please choose only one output option: --excel, --txt, or --json")
		os.Exit(1)
	}

	// Démarrer le crawling avec les paramètres choisis
	stats := startCrawling(siteURL, *excelFlag, *txtFlag, *jsonFlag, *verboseFlag)

	// Afficher le résumé des statistiques
	if stats != nil {
		fmt.Printf("Total: %d pages, %d inputs, %d buttons, %d hidden inputs\n",
			stats.TotalPages, stats.TotalInputs, stats.TotalButtons, stats.TotalHiddenInputs)
	}

	// Enregistrer les résultats en JSON si l'option est activée
	if *jsonFlag {
		elementsData := prepareElementData()
		saveResultsToJson("results.json", elementsData)
	}
}

func printUsage() {
	fmt.Println("Usage: go run main.go [--excel|--txt|--json|--verbose|--help] <url>")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

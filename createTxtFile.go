package main

import (
	"fmt"
	"os"
)

func saveResultsToTxt(stats *Stats) {
	file, err := os.Create("CrawlingResults.txt")
	if err != nil {
		fmt.Println("Failed to create text file:", err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("Total Pages: %d\n", stats.TotalPages))
	file.WriteString(fmt.Sprintf("Total Inputs: %d\n", stats.TotalInputs))
	file.WriteString(fmt.Sprintf("Total Hidden Inputs: %d\n", stats.TotalHiddenInputs))
	file.WriteString("URL\tInput ID\tInput Name\tInput Type\tHidden Inputs\n")
	for _, page := range visitedPages {
		for _, input := range page.Inputs {
			file.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\t%d\n", page.URL, input.ID, input.Name, input.Type, page.HiddenInputs))
		}
	}

	fmt.Println("Results saved to CrawlingResults.txt")
}

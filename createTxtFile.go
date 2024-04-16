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
	file.WriteString("URL\tInputs Count\tHidden Inputs\n")
	for _, pageInfo := range visitedPages {
		file.WriteString(fmt.Sprintf("%s\t%d\t%d\n", pageInfo.URL, pageInfo.InputsCount, pageInfo.HiddenInputs))
	}

	fmt.Println("Results saved to CrawlingResults.txt")
}

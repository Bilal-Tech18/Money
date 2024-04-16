package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// Fonction pour sauvegarder les résultats dans un fichier Excel
func saveResultsToExcel(stats *Stats) {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Total Pages")
	f.SetCellValue("Sheet1", "B1", stats.TotalPages)
	f.SetCellValue("Sheet1", "A2", "Total Inputs")
	f.SetCellValue("Sheet1", "B2", stats.TotalInputs)
	f.SetCellValue("Sheet1", "A3", "Total Hidden Inputs")
	f.SetCellValue("Sheet1", "B3", stats.TotalHiddenInputs)
	f.SetCellValue("Sheet1", "A5", "URL")
	f.SetCellValue("Sheet1", "B5", "Inputs Count")
	f.SetCellValue("Sheet1", "C5", "Hidden Inputs")

	row := 6 // Commencez à la ligne 6 pour les données des pages visitées
	for _, pageInfo := range visitedPages {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), pageInfo.URL)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), pageInfo.InputsCount)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), pageInfo.HiddenInputs)
		row++
	}

	// Sauvegardez le fichier au chemin spécifié
	filePath := "CrawlingResults.xlsx"
	if err := f.SaveAs(filePath); err != nil {
		fmt.Println("Failed to save Excel file:", err)
	} else {
		fmt.Println("Results saved to", filePath)
	}
}

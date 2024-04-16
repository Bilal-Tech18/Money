package main

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/xuri/excelize/v2"
)

func saveResultsToExcel(stats *Stats) {
    fmt.Println("Starting to save Excel file...") // Debug message
    // Obtenez le répertoire de travail courant
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println("Failed to get current working directory:", err)
        return
    }

    // Construisez le chemin complet du fichier
    filePath := filepath.Join(cwd, "CrawlingResults.xlsx")

    f := excelize.NewFile()
    f.SetCellValue("Sheet1", "A1", "Total Pages")
    f.SetCellValue("Sheet1", "B1", "Total Inputs")
    f.SetCellValue("Sheet1", "C1", "Total Hidden Inputs")
    f.SetCellValue("Sheet1", "A2", stats.TotalPages)
    f.SetCellValue("Sheet1", "B2", stats.TotalInputs)
    f.SetCellValue("Sheet1", "C2", stats.TotalHiddenInputs)

    // Sauvegardez le fichier au chemin spécifié
    if err := f.SaveAs(filePath); err != nil {
        fmt.Println("Failed to save Excel file:", err)
    } else {
        fmt.Println("Results saved to", filePath)
    }
}

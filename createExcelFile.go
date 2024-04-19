package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func saveResultsToExcel(stats *Stats) {
	f := excelize.NewFile()

	// Style pour les totaux
	totalStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#D3AF37"}, // Armenian Gold
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold: false,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if err != nil {
		fmt.Println("Error creating total style:", err)
		return
	}

	// Style pour les Hidden Inputs avec couleur de fond rouge
	HiddenStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#AA0505"}, // Rouge
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold: false,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if err != nil {
		fmt.Println("Error creating total style:", err)
		return
	}

	// Style pour la première ligne avec une couleur de fond dorée
	TopPageStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#B29700"}, // Light Gold
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold: false,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if err != nil {
		fmt.Println("Error creating total style:", err)
		return
	}

	// Style pour les lignes paires avec une couleur de fond bleue
	evenRowStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#2B6BBD"}, // Bleu
			Pattern: 1,
		},
	})
	if err != nil {
		fmt.Println("Error creating even row style:", err)
		return
	}

	// Style par défaut pour les lignes impaires
	oddRowStyle := 0

	// Application des styles pour les totaux
	f.SetCellStyle("Sheet1", "A1", "A1", totalStyle)
	f.SetCellStyle("Sheet1", "A2", "A2", totalStyle)
	f.SetCellStyle("Sheet1", "A3", "A3", totalStyle)

	// Ajout des données et des en-têtes
	f.SetCellValue("Sheet1", "A1", fmt.Sprintf("Total Pages = %d", stats.TotalPages))
	f.SetCellValue("Sheet1", "A2", fmt.Sprintf("Total Inputs = %d", stats.TotalInputs))
	f.SetCellValue("Sheet1", "A3", fmt.Sprintf("Total Hidden Inputs = %d", stats.TotalHiddenInputs))

	f.SetCellValue("Sheet1", "A5", "URL")
	f.SetCellValue("Sheet1", "B5", "Input ID")
	f.SetCellValue("Sheet1", "C5", "Input Name")
	f.SetCellValue("Sheet1", "D5", "Input Type")
	f.SetCellStyle("Sheet1", "A5", "E5", TopPageStyle)

	row := 6
	var lastURL string
	alternate := false

	for _, page := range visitedPages {
		firstInput := true // Utiliser cette variable pour déterminer si nous traitons le premier input d'une URL donnée
		for _, input := range page.Inputs {
			if page.URL != lastURL {
				alternate = !alternate
				lastURL = page.URL
				firstInput = true // Marquer que c'est le premier input de la nouvelle URL
			}

			styleToApply := oddRowStyle
			if alternate {
				styleToApply = evenRowStyle
			}

			// Si c'est le premier input de la nouvelle URL, imprimez l'URL
			if firstInput {
				f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), page.URL)
				firstInput = false
			} else {
				f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), "")
			}

			// Appliquez le style déterminé par alternate et si c'est une nouvelle URL
			f.SetCellStyle("Sheet1", fmt.Sprintf("A%d", row), fmt.Sprintf("E%d", row), styleToApply)

			// Définir les valeurs des autres cellules
			f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), input.ID)
			f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), input.Name)
			f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), input.Type)
			if input.Type == "hidden" {
				f.SetCellStyle("Sheet1", fmt.Sprintf("D%d", row), fmt.Sprintf("D%d", row), HiddenStyle)
			}

			row++ // Passez à la ligne suivante
		}
	}

	// Ajuster la largeur des colonnes
	f.SetColWidth("Sheet1", "A", "A", 150)
	f.SetColWidth("Sheet1", "B", "E", 35)

	// Sauvegarde du fichier
	if err := f.SaveAs("CrawlingResults.xlsx"); err != nil {
		fmt.Println("Failed to save Excel file:", err)
	}
}

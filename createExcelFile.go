package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func saveResultsToExcel(stats *Stats) {
	f := excelize.NewFile()

	totalStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#D3AF37"},
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

	TopPageStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#B29700"},
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

	evenRowStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#2B6BBD"}, 
			Pattern: 1,
		},
	})
	if err != nil {
		fmt.Println("Error creating even row style:", err)
		return
	}

	oddRowStyle := 0

	f.SetCellStyle("Sheet1", "A1", "A1", totalStyle)
	f.SetCellStyle("Sheet1", "A2", "A2", totalStyle)
	f.SetCellStyle("Sheet1", "A3", "A3", totalStyle)

	f.SetCellValue("Sheet1", "A1", fmt.Sprintf("Total Pages = %d", stats.TotalPages))
	f.SetCellValue("Sheet1", "A2", fmt.Sprintf("Total Inputs = %d", stats.TotalInputs))
	f.SetCellValue("Sheet1", "A3", fmt.Sprintf("Total Hidden Inputs = %d", stats.TotalHiddenInputs))

	f.SetCellValue("Sheet1", "A5", "URL")
	f.SetCellValue("Sheet1", "B5", "Input Name")
	f.SetCellValue("Sheet1", "C5", "Input Type")
	f.SetCellValue("Sheet1", "D5", "Input ID")
	f.SetCellStyle("Sheet1", "A5", "D5", TopPageStyle)

	row := 6
	var lastURL string
	alternate := false

	for _, page := range visitedPages {
		firstInput := true
		for _, input := range page.Inputs {
			if page.URL != lastURL {
				alternate = !alternate
				lastURL = page.URL
				firstInput = true
			}

			styleToApply := oddRowStyle
			if alternate {
				styleToApply = evenRowStyle
			}

			if firstInput {
				f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), page.URL)
				firstInput = false
			} else {
				f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), "")
			}

			f.SetCellStyle("Sheet1", fmt.Sprintf("A%d", row), fmt.Sprintf("D%d", row), styleToApply)

			f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), input.Name)
			f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), input.Type)
			f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), input.ID)
			if input.Type == "hidden" {
				f.SetCellStyle("Sheet1", fmt.Sprintf("C%d", row), fmt.Sprintf("C%d", row), HiddenStyle)
			}

			row++
		}
	}

	f.SetColWidth("Sheet1", "A", "A", 119)
	f.SetColWidth("Sheet1", "B", "D", 35)

	if err := f.SaveAs("CrawlingResults.xlsx"); err != nil {
		fmt.Println("Failed to save Excel file:", err)
	}
}

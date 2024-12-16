package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Properties struct {
	ElementType string `json:"elementType"`
	Placeholder string `json:"placeholder"`
	LinkTo      string `json:"linkTo"`
	Clickable   bool   `json:"clickable"`
	Visibility  bool   `json:"visibility"`
	ID          int    `json:"id"`
	Page        string `json:"page"`
	Redirection string `json:"redirection"`
	Type        string `json:"type"`
	Port        string `json:"port"`
	State       string `json:"state"`
	Service     string `json:"service"`
	Protocole   string `json:"protocole"`
}

func saveResultsToJson(filename string, elements []Properties) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(elements); err != nil {
		return fmt.Errorf("could not encode JSON: %v", err)
	}

	fmt.Println("Results saved to", filename)
	return nil
}

func convertToProperties(inputInfo InputInfo, pageURL string, elementType string, id int) Properties {
	return Properties{
		ElementType: elementType,
		Placeholder: inputInfo.Name,
		LinkTo:      "none",
		Clickable:   elementType == "button",
		Visibility:  true,
		ID:          id,
		Page:        pageURL,
		Redirection: "none",
		Type:        inputInfo.Type,
		Port:        "none",
		State:       "none",
		Service:     "none",
		Protocole:   "none",
	}
}

func prepareElementData() []Properties {
	var elementData []Properties
	idCounter := 1

	for _, page := range visitedPages {
		// Ajouter les inputs
		for _, input := range page.Inputs {
			prop := convertToProperties(input, page.URL, "input", idCounter)
			elementData = append(elementData, prop)
			idCounter++
		}

		// Ajouter les boutons
		for _, button := range page.Buttons {
			prop := convertToProperties(button, page.URL, "button", idCounter)
			elementData = append(elementData, prop)
			idCounter++
		}
	}

	return elementData
}

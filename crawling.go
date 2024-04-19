package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type InputInfo struct {
	ID   string
	Name string
	Type string
}

type PageInfo struct {
	URL          string
	Inputs       []InputInfo
	HiddenInputs int
}

type Stats struct {  // Définissez la structure Stats ici
	TotalPages        int
	TotalInputs       int
	TotalHiddenInputs int
}

var visitedPages []PageInfo // Stocke les informations sur chaque page visitée

// Fonction pour ajouter une page visitée à visitedPages
func addVisitedPage(url string, inputs []InputInfo, hiddenInputs int) {
	pageInfo := PageInfo{URL: url, Inputs: inputs, HiddenInputs: hiddenInputs}
	visitedPages = append(visitedPages, pageInfo)
}

// Fonction pour démarrer le crawling
func startCrawling(siteURL string, writeToExcel bool, verbose bool) *Stats {
	stats := &Stats{}
	c := colly.NewCollector(
		colly.AllowedDomains(getDomain(siteURL)),
	)

	// Visiter également les liens internes du domaine
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "/") {
			link = siteURL + link
		}
		c.Visit(link)
	})

	// Compter les éléments input et suivre les pages visitées
	c.OnHTML("html", func(e *colly.HTMLElement) {
		var inputs []InputInfo
		var hiddenInputCount int

		e.ForEach("input", func(_ int, el *colly.HTMLElement) {
			inputType := el.Attr("type")
			inputs = append(inputs, InputInfo{
				ID:   el.Attr("id"),
				Name: el.Attr("name"),
				Type: inputType,
			})
			if inputType == "hidden" {
				hiddenInputCount++
			}
		})

		if verbose {
			fmt.Println("Page visited:", e.Request.URL, "Total Inputs:", len(inputs), "Hidden Inputs:", hiddenInputCount)
		}
		addVisitedPage(e.Request.URL.String(), inputs, hiddenInputCount)
		stats.TotalPages++
		stats.TotalInputs += len(inputs)
		stats.TotalHiddenInputs += hiddenInputCount
	})

	// Gérer les erreurs
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// Visitez l'URL initiale
	c.Visit(siteURL)

	if writeToExcel {
		saveResultsToExcel(stats)
	} else {
		saveResultsToTxt(stats)
	}

	return stats
}

// Fonction pour extraire le domaine à partir de l'URL
func getDomain(url string) string {
	withoutProtocol := strings.TrimPrefix(url, "https://")
	parts := strings.Split(withoutProtocol, "/")
	return parts[0]
}

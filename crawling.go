package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

// Stats stocke les statistiques de crawling
type Stats struct {
	TotalPages        int
	TotalInputs       int
	TotalHiddenInputs int
}

// PageInfo stocke les informations sur chaque page visitée
type PageInfo struct {
	URL          string
	InputsCount  int
	HiddenInputs int
}

var visitedPages []PageInfo // Stocke les informations sur chaque page visitée

// Fonction pour ajouter une page visitée à visitedPages
func addVisitedPage(url string, inputsCount, hiddenInputs int) {
	pageInfo := PageInfo{URL: url, InputsCount: inputsCount, HiddenInputs: hiddenInputs}
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

	c.OnHTML("html", func(e *colly.HTMLElement) {
		var inputCount int
		var hiddenInputCount int

		// Compter tous les éléments input
		e.ForEach("input", func(_ int, el *colly.HTMLElement) {
			inputCount++
			stats.TotalInputs++
			// Compter spécifiquement les inputs de type hidden
			if el.Attr("type") == "hidden" {
				hiddenInputCount++
				stats.TotalHiddenInputs++
			}
		})

		// Enregistrez les informations sur la page visitée
		if verbose {
			fmt.Println("Page visited:", e.Request.URL, "Inputs:", inputCount, "Hidden Inputs:", hiddenInputCount)
		}
		addVisitedPage(e.Request.URL.String(), inputCount, hiddenInputCount)
		stats.TotalPages++
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

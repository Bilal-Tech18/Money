package main

import (
	"fmt"
	"net/url"
	"os"

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
	Buttons      []InputInfo
	HiddenInputs int
}

type Stats struct {
	TotalPages        int
	TotalInputs       int
	TotalButtons      int
	TotalHiddenInputs int
}

var visitedPages []PageInfo

func addVisitedPage(url string, inputs []InputInfo, buttons []InputInfo, hiddenInputs int) {
	pageInfo := PageInfo{URL: url, Inputs: inputs, Buttons: buttons, HiddenInputs: hiddenInputs}
	visitedPages = append(visitedPages, pageInfo)
}

func startCrawling(siteURL string, writeToExcel bool, writeToTxt bool, writeToJson bool, verbose bool) *Stats {
	stats := &Stats{}
	c := colly.NewCollector(
		colly.AllowedDomains(getDomain(siteURL)),
	)

	// Suivre les liens internes du domaine
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		parsedLink, err := url.Parse(link)
		if err != nil {
			return
		}
		if !parsedLink.IsAbs() {
			parsedLink = e.Request.URL.ResolveReference(parsedLink)
		}
		c.Visit(parsedLink.String())
	})

	// Compter les éléments input, button, textarea, et suivre les pages visitées
	c.OnHTML("html", func(e *colly.HTMLElement) {
		var inputs []InputInfo
		var buttons []InputInfo
		var hiddenInputCount int

		// Collecter les inputs
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

		// Collecter les textareas
		e.ForEach("textarea", func(_ int, el *colly.HTMLElement) {
			inputs = append(inputs, InputInfo{
				ID:   el.Attr("id"),
				Name: el.Attr("name"),
				Type: "textarea",
			})
		})

		// Collecter les boutons (balises <button>)
		e.ForEach("button", func(_ int, el *colly.HTMLElement) {
			buttons = append(buttons, InputInfo{
				ID:   el.Attr("id"),
				Name: el.Attr("name"),
				Type: "button",
			})
		})

		// Collecter les <input> de type bouton
		e.ForEach("input", func(_ int, el *colly.HTMLElement) {
			if el.Attr("type") == "button" || el.Attr("type") == "submit" {
				buttons = append(buttons, InputInfo{
					ID:   el.Attr("id"),
					Name: el.Attr("name"),
					Type: el.Attr("type"),
				})
			}
		})

		// Détecter les liens <a> ayant un rôle de bouton
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			if el.Attr("role") == "button" {
				buttons = append(buttons, InputInfo{
					ID:   el.Attr("id"),
					Name: el.Attr("name"),
					Type: "link-button",
				})
			}
		})

		// Détecter les <div> et <span> ayant un rôle de bouton ou des attributs d'accessibilité
		e.ForEach("div, span", func(_ int, el *colly.HTMLElement) {
			role := el.Attr("role")
			ariaPressed := el.Attr("aria-pressed")
			ariaHaspopup := el.Attr("aria-haspopup")

			// Vérifier les attributs qui pourraient indiquer un bouton
			if role == "button" || ariaPressed != "" || ariaHaspopup != "" {
				buttons = append(buttons, InputInfo{
					ID:   el.Attr("id"),
					Name: el.Attr("name"),
					Type: "div-button",
				})
			}
		})

		if verbose {
			fmt.Printf("Page visited: %s, Inputs: %d, Buttons: %d, Hidden Inputs: %d\n",
				e.Request.URL, len(inputs), len(buttons), hiddenInputCount)
			os.Stdout.Sync()
		}

		addVisitedPage(e.Request.URL.String(), inputs, buttons, hiddenInputCount)
		stats.TotalPages++
		stats.TotalInputs += len(inputs)
		stats.TotalButtons += len(buttons)
		stats.TotalHiddenInputs += hiddenInputCount
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.Visit(siteURL)

	if writeToExcel {
		saveResultsToExcel(stats)
	} else if writeToTxt {
		saveResultsToTxt(stats)
	} else if writeToJson {
		elementsData := prepareElementData()
		saveResultsToJson("results.json", elementsData)
	}

	return stats
}

func getDomain(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return ""
	}
	return u.Host
}

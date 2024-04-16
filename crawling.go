package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
    "net/url"
)

type Stats struct {
    TotalPages        int
    TotalInputs       int
    TotalHiddenInputs int
}

func startCrawling(siteURL string) *Stats {
    parsedURL, err := url.Parse(siteURL)
    if err != nil {
        fmt.Println("Could not parse URL:", err)
        return nil
    }
    domain := parsedURL.Hostname()

    c := colly.NewCollector(
        colly.AllowedDomains(domain),
    )

    stats := &Stats{} // Initialise stats une seule fois ici

    c.OnHTML("html", func(e *colly.HTMLElement) {
        var inputCount int
        var hiddenInputCount int

        // Compter tous les éléments input
        e.ForEach("input", func(_ int, el *colly.HTMLElement) {
            inputCount++
            stats.TotalInputs++ // Utiliser la même instance de Stats
            // Compter spécifiquement les inputs de type hidden
            if el.Attr("type") == "hidden" {
                hiddenInputCount++
                stats.TotalHiddenInputs++ // Utiliser la même instance de Stats
            }
        })

        // Afficher les résultats pour cette page
        fmt.Printf("Page visited: %s = %d inputs", e.Request.URL.String(), inputCount)
        stats.TotalPages++ // Incrémenter le total des pages visitées
        if hiddenInputCount > 0 {
            fmt.Printf(", %d inputs hidden", hiddenInputCount)
        }
        fmt.Println() // Nouvelle ligne pour séparer les entrées
    })

    c.OnError(func(_ *colly.Response, err error) {
        fmt.Println("Something went wrong:", err.Error())
    })

    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        c.Visit(e.Request.AbsoluteURL(link))
    })

    // Visiter l'URL initiale
    c.Visit(siteURL)
    return stats // Retourner la référence à stats
}


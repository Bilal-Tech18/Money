package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

func startCrawling(siteURL string) int {
    c := colly.NewCollector()

    var pageCount int // Compteur de pages

    // Définir ce qui se passe lorsque Colly visite une page
    c.OnHTML("html", func(e *colly.HTMLElement) {
        pageCount++
        fmt.Println("Visiting", e.Request.URL)
    })

    // Gérer les erreurs
    c.OnError(func(_ *colly.Response, err error) {
        fmt.Println("Something went wrong:", err.Error())
    })

    // Gérer les liens à visiter
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        // Visiter les liens
        c.Visit(e.Request.AbsoluteURL(link))
    })

    // Commencer le crawling à partir de l'URL fournie en argument
    c.Visit(siteURL)

    return pageCount // Retourner le nombre de pages visitées
}

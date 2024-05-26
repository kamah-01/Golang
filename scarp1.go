package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ain() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Link found:", link)
	})

	// Start scraping on https://httpbin.org
	c.Visit("https://httpbin.org")
}

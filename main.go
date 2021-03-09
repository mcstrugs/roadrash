package main

import (
    "fmt"
    "github.com/asciimoo/colly"
)

func main() {
    fmt.Println("Hello, World!")

    c := colly.NewCollector(
        colly.AllowedDomains("joemartinson.org"),
    )

    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://joemartinson.org/")
}

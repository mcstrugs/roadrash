package main

import (
    "fmt"
    "strings"
    "os"
    "github.com/asciimoo/colly"
)

func main() {
        data, _ := os.ReadFile("rx6000.txt")
        parsed_data := string(data)
        lines := strings.Split(parsed_data, "\n")
        for _, url  := range lines {
            checkStock(url)
        }
}

func checkStock(url string) bool {
    in_stock := true
    c := colly.NewCollector()

    c.OnHTML("span", func(e *colly.HTMLElement) {
        if(strings.Contains(e.Text, "OUT OF STOCK")) {
            in_stock = false
        }
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

    if(in_stock) {
        fmt.Println("IN STOCK")
    } else {
        fmt.Println("OUT OF STOCK")
    }

    return in_stock
}

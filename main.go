package main

import (
    "fmt"
    "strings"
    "github.com/asciimoo/colly"
)

func main() {
    checkStock("https://www.newegg.com/yeston-rx560d-4g-d5-ga/p/27N-0042-00028")
    checkStock("https://www.newegg.com/gigabyte-radeon-rx-6800-xt-gv-r68xtaorus-m-16gc/p/N82E16814932392")
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

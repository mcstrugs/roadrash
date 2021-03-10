package main

import (
    "fmt"
    "strings"
    "os"
    "github.com/asciimoo/colly"
)

func main() {
    checkList("rtx3000.txt")
    checkList("rx6000.txt")
}

func checkStock(url string) bool {
    in_stock := true
    c := colly.NewCollector()

    c.OnHTML("span", func(e *colly.HTMLElement) {
        if(strings.Contains(e.Text, "OUT OF STOCK")) {
            in_stock = false
        }
	})

	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL)
	//})

	c.Visit(url)

    if(in_stock) {
        fmt.Println("\x1B[32mIN STOCK: " + url + "\033[0m\t\t")
    } else {
        fmt.Println("OUT OF STOCK")
    }

    return in_stock
}

func linksFromFile(path string) []string {
    data, _ := os.ReadFile(path)
    parsed_data := string(data)
    return strings.Split(parsed_data, "\n")
}

func checkList(list string) {
    lines := linksFromFile(list)
    for _, url  := range lines {
        if (!(strings.Compare(url,"") == 0)) {
            checkStock(url)
        }
    }
}

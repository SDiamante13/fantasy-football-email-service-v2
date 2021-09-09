package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
	})

	c.OnHTML("table", func(element *colly.HTMLElement) {
		element.DOM.Find("thead tr th").Each(func(_ int, s *goquery.Selection) {
			fmt.Printf("header: %+v\n", s.Text())
		})

		element.DOM.Find("tbody tr").Each(func(_ int, s *goquery.Selection) {
			fmt.Printf("body: %+v\n", s.Text())
		})

	})

	c.Visit("https://www.fantasypros.com/nfl/rankings/waiver-wire-flex.php")
}

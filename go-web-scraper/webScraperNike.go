package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("div.product-price.us__styling.is--current-price.css-11s12ax[data-testid=product-price]", func(e *colly.HTMLElement) {
		price := e.Text
		fmt.Println("Product Price:", price)
	})

	// Visit the URL containing the target element
	err := c.Visit("https://www.nike.com/w/mens-shoes-nik1zy7ok")
	if err != nil {
		fmt.Println("Error visiting the website:", err)
	}
}

// web scraper with test site "scrapeMe" that uses pokemon for learning

package main 

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"encoding/json"
)

type pokemon struct {
	pokemon_name string
	pokemon_cost string
	pokemon_img string
}

func main() {

	c := colly.NewCollector(
		//Set allowable domains to only scrapeme.live test site
		colly.AllowedDomains("scrapeme.live/shop/", "www.scrapeme.com/shop/"),
	)

	//create a slice to store results in once scraped
	var pokemons []pokemon

	c.OnHTML("div.content div.primary", func(h *colly.HTMLElement) {
		fmt.Println("Raw HTML content", h.Text)
		//use declared struct to map scraping data to fields
		pokemon := pokemon{
			pokemon_name: h.ChildText("h2.woocommerce-loop-product__title"), 
			pokemon_cost: h.ChildText("span.price"), 
			pokemon_cost: h.ChildText("img", "src"),
		}

		pokemons = append(pokemon, pokemons)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	c.Visit("https://scrapeme.live/shop/")
	content, err := json.Marshal(pokemons)

	if err != nil {
		fmt.Println(err.Error())
	}

	//create json file from scraped data
	os.WriteFile("pokemon.json", content, 0644)
}
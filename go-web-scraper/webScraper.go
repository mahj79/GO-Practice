// webScraper.go

package main 

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"encoding/json"
)

type nfl_odds struct {
	home_team string `json:"home_team"`
	away_team string `json:"away_team"`
	game_odds string `json:"game_odds"`
}

func main() {

	c := colly.NewCollector(
		//Set allowable domains to only espn.com and www.espn.com
		colly.AllowedDomains("espn.com", "www.espn.com"),
	)

	//create a slice to store results in once scraped
	var nfl_odd []nfl_odds

	c.OnHTML("div.ScoreCell__Link__Event__Detail section.Card.gameModules.gameModules--mobile", func(h *colly.HTMLElement) {
		fmt.Println("Raw HTML content", h.Text)
		//use declared struct to map scraping data to fields
		nfl_odds := nfl_odds{
			home_team: h.ChildText("li.ScoreCell__Item.ScoreCell__Item--home"), 
			away_team: h.ChildText("li.ScoreCell__Item.ScoreCell__Item--away"), 
			game_odds: h.ChildText("div.ScoreCell__Odds"),
		}

		nfl_odd = append(nfl_odd, nfl_odds)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	c.Visit("https://www.espn.com/nfl/scoreboard")
	content, err := json.Marshal(nfl_odd)

	if err != nil {
		fmt.Println(err.Error())
	}

	//create json file from scraped data
	os.WriteFile("espn_nfl_odds.json", content, 0644)
}
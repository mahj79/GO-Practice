// webScraper.go

package main 

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type nfl_odds struct {
	game_date string
	home_team string
	away_team string
	game_time string
	game_odds string
}

func main() {
	fName := "nfl_odds.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	c := colly.NewCollector(
		//Set allowable domains to only espn.com and www.espn.com
		colly.AllowedDomains("espn.com", "www.espn.com"),
		//Cache responses to only perform one download of pages
		colly.CacheDir("./espn_cache"),
	)

	// Create another collector that is used to scrape details
	//detailCollector := c.Clone()

	spreads := make([]Spread, 0, 200)

	c.OnHtml("div[ScoreCell nfl ScoreCell--md ScoreCell--pre]", func(h *colly.HTMLElement)) {
		fmt.Println(h.Text)

	}

	//c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	c.Visit("https://www.espn.com/nfl/scoreboard")

}
package main

import (
	"fmt"
	"math/rand"
	"time"
)
func randomNumberGen(min float32, max float32) float32 {	
	return min + (max-min)*rand.Float32()
}
// Task implementation goes here
type Stock struct {
  name string
  price float32
}

func (stock *Stock) updateStock() {
  change := randomNumberGen(-10000, 10000)
  stock.price += change
}

func displayMarket(market []Stock) {
  for i:=0; i<len(market); i++{
    fmt.Println(market[i])
  }
}

func main() {
  rand.Seed(time.Now().UnixNano())
	// Function calls go here
  GOOG := Stock{"GOOG", 2313.50}
  AAPL := Stock{"AAPL", 157.28}
  FB := Stock{"FB", 203.77}
  TWTR := Stock{"TWTR", 50}
  stockMarket := []Stock{GOOG, AAPL, FB, TWTR}
  displayMarket(stockMarket)
  stockMarket[0].updateStock()
  stockMarket[1].updateStock()
  stockMarket[2].updateStock()
  stockMarket[3].updateStock()
  displayMarket(stockMarket)
}
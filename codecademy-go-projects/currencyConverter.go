package main

import (
	"fmt"
)

func main() {
  var dollarAmount float32;
  var currency string;
  currencies := map[string]float32{
    "JPY": 130.2,
    "EUR": 0.95,
  }
  fmt.Println("Please enter a dollar amount:")
  fmt.Scan(&dollarAmount)
  if dollarAmount < 0 {
    fmt.Println("Please enter a valid number")
  } else {
    fmt.Println("What currency would you like to select?")
    fmt.Scan(&currency)
  }
  rate, isValid := currencies[currency]
  if isValid {
    rate *= dollarAmount
    fmt.Println("Your money is worth", rate, "after converting to the selected currency.")
  } else {
    fmt.Println("That currency is not available.")
  }

}
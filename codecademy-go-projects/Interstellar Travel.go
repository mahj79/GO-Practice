package main

import "fmt"

func fuelGauge(fuel int) {
  fmt.Println("You have", fuel, "gallons of fuel left!")
}

func calculateFuel(planet string) int {
  var fuel int
  if (planet == "Venus") {
    fuel = 300000
  } else if (planet == "Mercury") {
    fuel = 500000
  } else if (planet == "Mars") {
    fuel = 700000
  } else {
    fuel = 0
  }

  return fuel
}

func greetPlanet(planet string) {
  fmt.Println("Welcome! You are on your way to", planet)
}

func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining int = fuel
  var fuelCost int = calculateFuel(planet)

  if (fuelRemaining >= fuelCost) {
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }

  return fuelRemaining
}

func main() {
  var fuel int = 1000000
  var planetChoice = "Venus"

  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
  
}
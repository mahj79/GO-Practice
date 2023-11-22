package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano());
  var isHeistOn = true;
  var eludedGuards = rand.Intn(100);

  if (eludedGuards >= 50) {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    fmt.Println("Plan a better disguise next time?")
  }

  var openedVault = rand.Intn(100);
  if (eludedGuards >= 50 && openedVault >= 70 && isHeistOn == true) {
    fmt.Println("Grab and GO!");
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("The vault won't open!")
  }

  var leftSafely = rand.Intn(5);
  if (isHeistOn == true) {
    switch leftSafely {
      case 0: 
      isHeistOn = false
      fmt.Println("Looks like you messed up.")
      case 1: 
      isHeistOn = false
      fmt.Println("You tripped the alarm!")
      case 2: 
      isHeistOn = false
      fmt.Println("Looks like the security guard was faster!")
      case 3: 
      isHeistOn = false
      fmt.Println("Should've remembered the get away car....")
      default:
      isHeistOn = true
      fmt.Println("Start the getaway car!")
    }
  }

  if (isHeistOn == true) {
    var amtStolen = 10000 + rand.Intn(1000000)
    fmt.Println(amtStolen)
  }

  fmt.Println(isHeistOn)
}

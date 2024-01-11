package main

import (
  "fmt"
  "math/rand"
  "time"
)

func getRandomElement(slice []string) string{
  length := len(slice);
  randomValue := rand.Intn(length);
  return slice[randomValue]
}

func main() {
  rand.Seed(time.Now().UnixNano())
  guestList := [4]string{"Jordyn", "Jack", "Cleo", "Nug"}
  catStorage := [4]string{"Toy Chest", "Crate", "Box", "Dryer"}
  fmt.Println(guestList, catStorage)
  guestListSlice := guestList[:];
  catStorageSlice := catStorage[:];
  randomGuest := getRandomElement(guestListSlice);
  randomObject := getRandomElement(catStorageSlice);

  fmt.Println("I can't believe", randomGuest, "hid the cat in the", randomObject)
}
//adding account struct for simulating account at a bank that sends json out

package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type Account struct {
	ID int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Number int64 `json:"number"`
	Balance int64 `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount (FirstName, LastName string) *Account {
	return &Account {
		FirstName: FirstName,
		LastName: LastName,
		Number: int64(rand.Intn(10000000)),
		CreatedAt: time.Now().UTC(),
	}
}
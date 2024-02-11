//project to build a go crud api

package main

import (
	"flag"
	"log"
	"fmt"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account")

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "jack", "maho", "panther9999")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	
	if *seed {
		fmt.Println("seeding the database")
		seedAccounts(store)
	}
	
	server := NewAPIServer(":3000", store)
	server.Run()
}
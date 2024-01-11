// go practice building a web server

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Building a web server!")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World\n")
		io.WriteString(w, r.Method)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
package main

import (
	"github.com/akaKAIN/go-vendor-tests/v2/getter"
	"log"
)

func main() {
	url := "https://github.com/akaKAIN/go-vendor-tests/v2"

	code, err := getter.GetStatusCode(url)
	if err != nil {
		log.Println(err)
	}

	log.Println("Code:", code)

}

package main

import "github.com/akaKAIN/go-gb-level-2/myatomic"

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	count := 1000
	myatomic.StartGo(count, myatomic.SimpleHandler)


}

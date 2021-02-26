package main

import (
	"fmt"
	"time"
)

func main() {
	var count int
	go ReadChan(&count)
	time.Sleep(time.Second)
	count = 20
	time.Sleep(10 * time.Second)
}

func ReadChan(num *int) {
	fmt.Println(*num)
	time.Sleep(2 * time.Second)
	fmt.Println(*num)
}

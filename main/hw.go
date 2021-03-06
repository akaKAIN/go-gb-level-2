package main

import (
	"fmt"
	"github.com/akaKAIN/go-gb-level-2/myatomic"
)

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	// Task 1
	count := 1000
	myatomic.StartGo(count, myatomic.SimpleHandler)

	// Task 2
	arr := myatomic.NewMutexIntArray(1, 2, 3, 4)
	arr.Push(5)
	old, err := arr.Replace(2, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(old, arr.ArrayBody)

	if _, err = arr.Replace(10, 11); err != nil {
		fmt.Println(err)
	}

	// Task 3
	myatomic.StartReadAndWrite(900, 100)

}

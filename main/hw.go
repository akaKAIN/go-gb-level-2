package main

import (
	"fmt"
	"github.com/akaKAIN/go-gb-level-2/myatomic"
)

func main() {
	// Task 1
	count := 1000
	myatomic.StartGo(count, myatomic.SimpleWorker)

	// Task 2
	arr := myatomic.NewMutexIntArray(1, 2, 3, 4)
	arr.Push(5)
	old, err := arr.Replace(2, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(old, arr.ArrayBody)

	if _, err = arr.Replace(10, 11); err != nil {
		fmt.Println(err)
		return
	}

	// Task 3
	store := &myatomic.IntMapRWM{
		Map: make(map[int]int),
	}
	myatomic.Fill(store, 1000)
	myatomic.StartReadAndWrite(900, 100, store)

}

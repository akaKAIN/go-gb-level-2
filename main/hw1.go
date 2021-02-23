package main

import (
	"fmt"
	"github.com/akaKAIN/go-gb-level-2/myerrors"
	"github.com/akaKAIN/go-gb-level-2/myfiles"
	"log"
)

func main() {
	CallPanic()

	err := myfiles.WriteNewFile("testFile.txt", []byte("Test"))
	if err != nil {
		log.Println("Error: ", err)
	}
}

func CallPanic() {
	defer recoverIt()

	var result int
	for i := 10; i > -5; i-- {
		result = 10 / i
	}
	fmt.Println(result)
}

func recoverIt() {
	if r := recover(); r != nil {
		err := fmt.Sprintf("%s", r)
		fmt.Println(myerrors.NewErrorTime(err))
	}
}

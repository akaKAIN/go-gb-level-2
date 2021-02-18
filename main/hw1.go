package main

import (
	"fmt"
	"os"
)

func main() {
	var err *error
	defer fmt.Println("Error: %s\n", err)

	_, err2 := os.Open("file")
	fmt.Println(err2)

	err = &err2
}

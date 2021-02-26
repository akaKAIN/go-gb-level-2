package demo

import "fmt"

func ExampleUser_String() {
	name := "Ivan"
	age := 22
	user := User{
		Name: name,
		Age:  age,
	}
	fmt.Println(user)
	// Output: user: Ivan, 22 years old
}

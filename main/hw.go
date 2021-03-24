package main

import (
	"fmt"
	"github.com/akaKAIN/go-gb-level-2/myreflect"
	"log"
)


/*
   1. Написать функцию, которая принимает на вход структуру in
   (struct или кастомную struct) и values map[string]interface{}
   (key - название поля структуры, которому нужно присвоить value этой мапы).

   Необходимо по значениям из мапы изменить входящую структуру in с помощью пакета reflect.
   Функция может возвращать только ошибку error.
   Написать к данной функции тесты (чем больше, тем лучше - зачтется в плюс).

   2. Написать функцию, которая принимает на вход имя файла и название функции.
   Необходимо подсчитать в этой функции количество вызовов асинхронных функций.
   Результат работы должен возвращать количество вызовов int и ошибку error.
   Разрешается использовать только go/parser, go/ast и go/token.

   3*. Написать кодогенератор под какую-нибудь задачу.

*/

//go:generate go run hw.go
func main() {
	TaskOne()
	TaskTwo()
}

func TaskOne() {
	var (
		name = "Ivan"
		age uint8 = 13
	)
	user := new(myreflect.User)
	m := map[string]interface{}{
		"Name": name,
		"Age":  age,
	}
	if err := myreflect.UpdateStruct(user, m); err != nil {
		log.Println(err)
	}
}

func TaskTwo()  {
	n, err := myreflect.Analyze("workers/example.go", "SoftShotDown")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

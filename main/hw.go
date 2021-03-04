package main

import (
	"github.com/akaKAIN/go-gb-level-2/myatomic"
)

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	//db, err := mydb.DbMutex()
	//if err != nil {
	//	log.Println(err)
	//}
	//var book Book
	//r := db.QueryRow("SELECT * from book limit 1")
	//err = r.Scan(&book.id, &book.name, &book.author)
	//fmt.Println("OUTPUT", book, err)

	myatomic.SimpleCounter()
}

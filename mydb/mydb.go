package mydb

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"sync"
)

var (
	lock sync.Mutex
	db   *sql.DB
)

func DbMutex() (*sql.DB, error) {
	lock.Lock()
	defer lock.Unlock()

	if db != nil {
		return db, nil
	}

	var err error
	db, err = sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

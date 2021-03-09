package main

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"sync"
)

var (
	lock sync.Mutex
	db   *sql.DB
)

/*
func main() {
	db, err := DbMutex()
	if err != nil {
		panic(err)
	}

	var v int
	r := db.QueryRow("SELECT 1")
	err = r.Scan(&v)
	fmt.Println(v, err)
}
*/
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

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "test_db"
	tablename = "test_table"
)

func main () {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connInfo)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected!")

	initTable(db)
	defer deleteTable(db)

	fmt.Println(getValidation(db, "Элемент данных 1",
		time.Date(2021, time.September, 29, 21,00,0,0,time.UTC),
		time.Date(2021, time.September, 29, 22,30,0,0,time.UTC)))

	fmt.Println(getValidationSQL(db, "Элемент данных 1",
		time.Date(2021, time.September, 29, 21,00,0,0,time.UTC),
		time.Date(2021, time.September, 29, 22,30,0,0,time.UTC)))
	// ---
	fmt.Println(getValidation(db, "Элемент данных 2",
		time.Date(2021, time.September, 30, 15,0,0,0,time.Local),
		time.Date(2021, time.September, 30, 17,0,0,0,time.Local)))

	fmt.Println(getValidationSQL(db, "Элемент данных 2",
		time.Date(2021, time.September, 30, 15,0,0,0,time.Local),
		time.Date(2021, time.September, 30, 17,0,0,0,time.Local)))

}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"
)


type testpair struct {
	elem string
	timerange
	result bool
}

var tests = []testpair {
	// Пересечение
	{"Элемент данных 2",
		timerange{time.Date(2021, time.September, 30, 15,0,0,0,time.UTC),
					time.Date(2021, time.September, 30, 17,0,0,0,time.UTC)}, false},
	// Сопрекасновение диапазонов
	{"Элемент данных 2",
		timerange{time.Date(2021, time.September, 30, 16,0,0,0,time.UTC),
					time.Date(2021, time.September, 30, 17,0,0,0,time.UTC)}, false},
	// Пересечение с несколькими
	{"Элемент данных 2",
		timerange{time.Date(2021, time.September, 28, 11,0,0,0,time.UTC),
					time.Date(2021, time.September, 30, 17,0,0,0,time.UTC)}, false},
    // Между двумя элементами
	{"Элемент данных 2",
		timerange{time.Date(2021, time.September, 29, 22,0,0,0,time.UTC),
					time.Date(2021, time.September, 29, 23,0,0,0,time.UTC)}, true},
	// После
	{"Элемент данных 1",
		timerange{time.Date(2021, time.October, 2, 15,0,0,0,time.UTC),
					time.Date(2021, time.October, 2, 17,0,0,0,time.UTC)}, true},
		// Пересечение
	{"Элемент данных 1",
		timerange{time.Date(2021, time.September, 29, 22,0,0,0,time.UTC),
					time.Date(2021, time.September, 29, 22,30,0,0,time.UTC)}, false},
	// Пересечение
	{"Элемент данных 1",
		timerange{time.Date(2021, time.September, 29, 22,10,0,0,time.UTC),
					time.Date(2021, time.September, 29, 22,30,0,0,time.UTC)}, false},
	// Пересечение
	{"Элемент данных 1",
		timerange{time.Date(2021, time.September, 29, 21,30,0,0,time.UTC),
					time.Date(2021, time.September, 29, 22,30,0,0,time.UTC)}, false},
}

func TestGetValidation(t *testing.T) {

	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	initTable(db)
	defer deleteTable(db)

	for i, pair := range tests {
		v := getValidation(db, pair.elem, pair.dateStart, pair.dateEnd)
		if v != pair.result {
			t.Error(
				"Тест №", i,
				"Для:", pair.elem, pair.dateStart, pair.dateEnd,
				"ожидалось:", pair.result,
				", получили:", v,
			)
		}
	}
}

func TestGetValidationSQL(t *testing.T) {

	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	initTable(db)
	defer deleteTable(db)

	for i, pair := range tests {
		v := getValidationSQL(db, pair.elem, pair.dateStart, pair.dateEnd)
		if v != pair.result {
			t.Error(
				"Тест №", i,
				"Для:", pair.elem, pair.dateStart, pair.dateEnd,
				"ожидалось:", pair.result,
				", получили:", v,
			)
		}
	}
}
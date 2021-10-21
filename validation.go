package main

import (
	"database/sql"
	"log"
	"time"
)

type timerange struct {
	dateStart time.Time
	dateEnd time.Time
}

func getValidation (db *sql.DB, elem string, dateStart time.Time, dateEnd time.Time) bool {
	result, err := db.Query("select start_time, end_time from " + tablename + " where tech_window = '" + elem + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	for result.Next() {
		r := timerange{}
		err = result.Scan(&r.dateStart, &r.dateEnd)
		if err != nil {
			log.Fatal(err)
			continue
		}

		if (r.dateEnd.After(dateStart) || r.dateEnd.Equal(dateStart)) && (r.dateStart.Before(dateEnd) || r.dateStart.Equal(dateEnd)) {
			return false
		}
	}

	return true
}

func getValidationSQL (db *sql.DB, elem string, dateStart time.Time, dateEnd time.Time) bool {
	result, err := db.Query("select every(isempty(tstzrange (start_time, end_time, '[]') * " +
		"tstzrange (to_timestamp('" + dateStart.Format("02-01-2006 15:04") + "', 'dd-mm-yyyy hh24:mi'), " +
		"to_timestamp('" + dateEnd.Format("02-01-2006 15:04") + "', 'dd-mm-yyyy hh24:mi'),'[]'))) " +
		"from " + tablename + " " +
		"where tech_window = '" + elem + "'")
	if err != nil{
		log.Fatal(err)
	}
	defer result.Close()

	var res bool

	for result.Next() {
		result.Scan(&res)
	}

	return res
}
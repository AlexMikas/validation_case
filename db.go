package main

import (
	"database/sql"
	"log"
)

func initTable(db *sql.DB)  {
	_, err := db.Exec("create table if not exists test_table ( tech_window varchar(255), start_time timestamp, end_time timestamp )")

	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("insert into test_table (tech_window, start_time, end_time) values ('Элемент данных 1', to_timestamp('27-09-2021 20:00', 'dd-mm-yyyy hh24:mi'), to_timestamp('27-09-2021 23:00', 'dd-mm-yyyy hh24:mi'))")
	_, err = db.Exec("insert into test_table (tech_window, start_time, end_time) values ('Элемент данных 2', to_timestamp('28-09-2021 12:00', 'dd-mm-yyyy hh24:mi'), to_timestamp('28-09-2021 16:00', 'dd-mm-yyyy hh24:mi'))")
	_, err = db.Exec("insert into test_table (tech_window, start_time, end_time) values ('Элемент данных 1', to_timestamp('29-09-2021 22:00', 'dd-mm-yyyy hh24:mi'), to_timestamp('29-09-2021 23:00', 'dd-mm-yyyy hh24:mi'))")
	_, err = db.Exec("insert into test_table (tech_window, start_time, end_time) values ('Элемент данных 2', to_timestamp('30-09-2021 12:00', 'dd-mm-yyyy hh24:mi'), to_timestamp('30-09-2021 16:00', 'dd-mm-yyyy hh24:mi'))")
	_, err = db.Exec("insert into test_table (tech_window, start_time, end_time) values ('Элемент данных 3', to_timestamp('01-10-2021 20:00', 'dd-mm-yyyy hh24:mi'), to_timestamp('01-10-2021 23:00', 'dd-mm-yyyy hh24:mi'))")

	if err != nil {
		log.Fatal(err)
	}
}

func deleteTable(db *sql.DB)  {
	_, err := db.Exec("drop table " + tablename)

	if err != nil {
		log.Fatal(err)
	}
}

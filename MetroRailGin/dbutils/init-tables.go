package dbutils

import (
	"database/sql"
	"fmt"
	"log"
)

func Initialize(db *sql.DB) {
	statement, err := db.Prepare(train)
	if err != nil {
		log.Println(err)
	}

	//Create train table
	_, statementerror := statement.Exec()
	if statementerror != nil {
		log.Println("Table already exists")
	}
	statement1, err0 := db.Prepare(station)
	if err0 != nil {
		log.Println(err0)
		fmt.Println(err0)
	}
	_, aerr := statement1.Exec()
	if aerr != nil {
		log.Println(aerr)
	}
	statement2, err1 := db.Prepare(schedule)
	if err1 != nil {
		log.Println(err1)
		fmt.Println(err1)
	}
	_, berr := statement2.Exec()
	if berr != nil {
		log.Println(berr)
		fmt.Println(berr)
	}
	log.Println("All tables created/initialized successfully!")
}

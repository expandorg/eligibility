package main

import (
	"fmt"
	"log"

	"github.com/gemsorg/eligibility/pkg/database"
	"github.com/jmoiron/sqlx"
	env "github.com/joho/godotenv"
)

func main() {
	err := env.Load()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("mysql connection error", err)
	}
	defer db.Close()

	err = seed(db)
	if err != nil {
		log.Fatalln(err)
	}
}

func seed(db *sqlx.DB) error {
	fmt.Println("seeding db")
	fs := `INSERT INTO filters (type, value)
		VALUES
			('Education','associate degree'),
			('Education','bachelor degree'),
			('Education','doctorate degree'),
			('Education','high school graduate, diploma or the equivalent (for example: GED)'),
			('Education','master degree'),
			('Education','no schooling completed'),
			('Education','nursery school to 8th grade'),
			('Education','professional degree'),
			('Education','some college credit, no degree'),
			('Education','some high school, no diploma'),
			('Education','trade/technical/vocational training'),
			('Gender','female'),
			('Gender','male'),
			('Interest','ai'),
			('Interest','machine learning'),
			('Language','arabic'),
			('Language','chinese'),
			('Language','english'),
			('Language','french'),
			('Language','german'),
			('Language','japanese'),
			('Language','portuguese'),
			('Language','russian'),
			('Language','spanish'),
			('Availability','1 - 5 hours'),
			('Availability','5  - 10 hours'),
			('Availability','10 - 20 hours'),
			('Availability','20 - 40 hours'),
			('Availability','40 hours+')`
	_, err := db.Exec(fs)
	if err != nil {
		fmt.Println("error seeding", err)
		return err
	}
	return nil
}

package database

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const connMaxLifetime = time.Second * 5

func Connect() (*sqlx.DB, error) {
	host := os.Getenv("ELIG_DB_HOST")
	port := os.Getenv("ELIG_DB_PORT")
	name := os.Getenv("ELIG_DB")
	user := os.Getenv("ELIG_DB_USER")
	password := os.Getenv("ELIG_DB_PASSWORD")
	if password != "" {
		password = ":" + password
	}

	connection := fmt.Sprintf(`%s%s@tcp([%s]:%s)/%s?parseTime=true`, user, password, host, port, name)
	db, err := sqlx.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(connMaxLifetime)

	for {
		err = db.Ping()
		if err == nil {
			break
		}
		fmt.Println("Retrying connection:", err)
		time.Sleep(time.Second)
	}
	return db, nil
}

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
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	if password != "" {
		password = ":" + password
	}

	connection := fmt.Sprintf(`%s%s@tcp([%s]:3306)/%s?parseTime=true`, user, password, host, name)
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

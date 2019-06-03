package test

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func Setup() *sqlx.DB {
	mockDB, _, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return sqlxDB
}

type fakeDB struct{}

func (db *fakeDB) Select(dest interface{}, query string, args ...interface{}) error {
	return nil
}

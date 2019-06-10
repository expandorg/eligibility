package test

import (
	"database/sql"

	"github.com/gemsorg/eligibility/mock"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func Setup() (*sql.DB, *sqlx.DB, sqlmock.Sqlmock) {
	return mock.Mysql()
}

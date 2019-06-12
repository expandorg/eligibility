package test

import (
	"database/sql"

	"github.com/gemsorg/eligibility/pkg/mock"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

const bearer = "Bearer 123"

func Setup() (*sql.DB, *sqlx.DB, sqlmock.Sqlmock) {
	return mock.Mysql()
}

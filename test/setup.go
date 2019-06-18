package test

import (
	"database/sql"
	"os"

	"github.com/gemsorg/eligibility/pkg/mock"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

var jwtToken, _ = mock.GenerateJWT(8)

var bearer = "Bearer " + jwtToken

func Setup() (*sql.DB, *sqlx.DB, sqlmock.Sqlmock) {
	os.Setenv("JWT_SECRET", mock.JWT_SECRET)
	return mock.Mysql()
}

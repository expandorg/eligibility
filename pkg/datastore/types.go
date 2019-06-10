package datastore

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

type Time struct {
	Time  time.Time
	Valid bool
}

func (n *Time) Scan(value interface{}) error {
	if value == nil {
		n.Time = time.Time{}
		n.Valid = false
		return nil
	}
	timeValue, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("null: cannot scan type %T into null.Time: %v", value, value)
	}
	n.Time = timeValue
	n.Valid = true
	return nil
}

func (n Time) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}

func NewTime(n interface{}) Time {
	if n == nil {
		return Time{
			Valid: false,
		}
	}
	return Time{
		Time:  n.(time.Time),
		Valid: true,
	}
}

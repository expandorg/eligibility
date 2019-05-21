package datastore

import "fmt"

type FilterNotFound struct {
	ids []int
}

func (err FilterNotFound) Error() string {
	return fmt.Sprintf("One or more of these attributes don't exist: %v", err.ids)
}

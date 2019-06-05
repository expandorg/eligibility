package service

import (
	"fmt"
)

type ProfileNotFound struct {
	WorkerID string
}

func (err *ProfileNotFound) Error() string {
	return fmt.Sprintf("Couldn't find a profile for worker_id: %s", err.WorkerID)
}

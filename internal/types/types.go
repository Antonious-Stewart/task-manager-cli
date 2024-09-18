package types

import "time"

type Task struct {
	ID          int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      string
}

type Status int

const (
	TODO Status = iota
	IN_PROGRESS
	DONE
)

func (s Status) String() string {
	return [...]string{"todo", "in-progress", "done"}[s]
}

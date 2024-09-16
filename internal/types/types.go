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
	switch s {
	case TODO:
		return "todo"
	case IN_PROGRESS:
		return "in_progress"
	case DONE:
		return "done"
	default:
		return "unknown"
	}
}

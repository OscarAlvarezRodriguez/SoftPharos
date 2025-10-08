package role

import "time"

type Role struct {
	ID          int
	Name        string
	Description *string
	CreatedAt   time.Time
}

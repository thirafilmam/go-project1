package entities

import "time"

type Category struct {
	ID         int
	Name       string
	Created_at time.Time
	Updated_at time.Time
}

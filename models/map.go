package models

import "time"

// Map is the model of key-value data
type Map struct {
	ID        int
	Key       string `sql:"size:255"`
	Value     string
	CreatedAt time.Time
}

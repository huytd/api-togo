package models

import "time"

// Map is the model of key-value data
type Map struct {
	ID        int       `json:"id"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created"`
}

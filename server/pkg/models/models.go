package models

import (
	"time"
)

type User struct {
	ID      int
	Name    string
	Mobile  string
	Created time.Time
}

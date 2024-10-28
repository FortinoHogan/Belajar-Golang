package entities

import "time"

type Product struct {
	Id          uint
	Name        string
	Category    Category
	Stock       int64
	Description string
	Created_at  time.Time
	Updated_at  time.Time
}

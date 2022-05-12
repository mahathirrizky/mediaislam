package videomateri

import (
	"mediaislam/submateri"
	"time"
)

type VideomateriTable struct {
	ID          int
	SubmateriID int
	Link        string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	SUbmateri   submateri.SubmateriTable
}

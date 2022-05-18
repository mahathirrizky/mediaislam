package videomateri

import (
	"time"
)

type VideomateriTable struct {
	ID          int
	SubmateriID int `gorm:"foreignkey:SubmateriRefer"`
	Link        string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

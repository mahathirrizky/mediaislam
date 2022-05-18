package submateri

import (
	"mediaislam/videomateri"
	"time"
)

type SubmateriTable struct {
	ID          int
	MateriID    int
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Videomateri []videomateri.VideomateriTable `gorm:"foreignKey:SubmateriID;references:ID"`
}

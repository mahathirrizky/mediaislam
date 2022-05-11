package submateri

import (
	"mediaislam/materi"
	"time"
)

type SubmateriTable struct {
	ID          int
	MateriID    int
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Materi      materi.MateriTable
}

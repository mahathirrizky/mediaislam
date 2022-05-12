package submateri

import (
	"time"
)

type SubmateriTable struct {
	ID        int
	MateriID  int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

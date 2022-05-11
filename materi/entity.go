package materi

import (
	"mediaislam/ustadz"
	"time"
)

type MateriTable struct {
	ID          int
	UstadzID    int
	Name        string
	Description string
	Slug        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      int
	Ustadz      ustadz.UstadzTable
}


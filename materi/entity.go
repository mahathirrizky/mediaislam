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

type SubmateriTable struct {
	ID          int
	MateriID    int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type VideomateriTable struct {
	ID          int
	SubmateriID int
	Link        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

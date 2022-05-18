package materi

import (
	"mediaislam/submateri"
	"mediaislam/user"
	"mediaislam/ustadz"
	"time"
)

type MateriTable struct {
	ID            int
	UstadzID      int
	UserID        int
	Name          string
	Description   string
	Slug          string
	ImageFileName string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Ustadz        ustadz.UstadzTable
	User          user.UserTable
	Submateri     []submateri.SubmateriTable `gorm:"foreignKey:MateriID;references:ID"`
}


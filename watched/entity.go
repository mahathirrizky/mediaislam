package watched

import (
	"mediaislam/user"
	"mediaislam/videomateri"
	"time"
)

type WatchedTable struct {
	ID            int `json:"id"`
	UserID        int `json:"user_id"`
	VideomateriID int `json:"videomateri_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	User          user.UserTable
	Videomateri   videomateri.VideomateriTable
}

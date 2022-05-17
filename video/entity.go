package video

import (
	"mediaislam/user"
	"time"
)

type VideoTable struct {
	ID            int
	UserID        int
	Name          string
	Description   string
	Link          string
	ImageFileName string
	Type          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	User          user.UserTable
}

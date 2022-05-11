package subscribe

import (
	"mediaislam/materi"
	"mediaislam/user"
	"time"
)

type SubscribeTable struct {
	ID        int  `json:"id"`
	UserID    int  `json:"user_id"`
	MateriID  int  `json:"materi_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.UserTable
	Materi    materi.MateriTable
}

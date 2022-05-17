package subscribe

import (
	"mediaislam/materi"
	"mediaislam/user"
	"time"
)

type SubscribeTable struct {
	ID        int  
	UserID    int  
	MateriID  int  
	AvatarFileName string 
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.UserTable
	Materi    materi.MateriTable
}

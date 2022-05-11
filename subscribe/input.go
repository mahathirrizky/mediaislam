package subscribe

import (
	"mediaislam/user"
)

type CreateSubscribeInput struct {
	MateriID int `json:"materi_id" binding:"required"`
	User     user.UserTable
}

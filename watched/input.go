package watched

import "mediaislam/user"

type CreateWatchedInput struct {
	VideomateriID int `json:"videomateri_id" binding:"required"`
	User          user.UserTable
}

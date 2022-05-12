package materi

import "mediaislam/user"

type GetMateriDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateMateriInput struct {
	UstadzID    int    `json:"ustadz_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	User        user.UserTable
}

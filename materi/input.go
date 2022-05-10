package materi

import "mediaislam/user"

type GetMateriDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetSubmateriDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetVideomateriDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateUstadzInput struct {
	Name string `json:"name" binding:"required"`
	User user.UserTable
}

type CreateMateriInput struct {
	UstadzID    int    `json:"ustadz_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	User        user.UserTable
}

type CreateSubmateriInput struct {
	MateriID    int    `json:"materi_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	User        user.UserTable
}

type CreateVideomateriInput struct {
	SubmateriID int    `json:"submateri_id" binding:"required"`
	Link        string `json:"link" binding:"required"`
	Description string `json:"description" binding:"required"`
	User        user.UserTable
}

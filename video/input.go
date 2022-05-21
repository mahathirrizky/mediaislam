package video

import "mediaislam/user"

type CreateVideoInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Link        string `json:"link" binding:"required"`
	Type        string `json:"type" binding:"required"`
	User        user.UserTable
}

type GetVideoDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetVideoTematikInput struct {
	Type string `json:"type" binding:"required"`
}

package submateri

type CreateSubmateriInput struct {
	MateriID int    `json:"materi_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type GetSubmateriDetailInput struct {
	ID int `json:"id" binding:"required"`
}

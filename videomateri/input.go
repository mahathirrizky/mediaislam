package videomateri

type CreateVideomateriInput struct {
	SubmateriID int    `json:"submateri_id" binding:"required"`
	Link        string `json:"link" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type GetVideomateriDetailInput struct {
	ID int `json:"id" binding:"required"`
}
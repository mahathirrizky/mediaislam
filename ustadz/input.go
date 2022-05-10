package ustadz

type RegisterUstadzInput struct {
	Name string `json:"name" binding:"required"`
}

type GetUstadzDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
package submateri

type CreateSubmateriFormatter struct {
	ID       int    `json:"id"`
	MateriID int    `json:"materi_id"`
	Name     string `json:"name"`
}

func FormatCreateSubmateri(submateri SubmateriTable) CreateSubmateriFormatter {
	createSubmateriFormatter := CreateSubmateriFormatter{
		ID:       submateri.ID,
		MateriID: submateri.MateriID,
		Name:     submateri.Name,
	}
	return createSubmateriFormatter
}

type UpdateSubmateriFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatUpdateSubmateri(submateri SubmateriTable) UpdateSubmateriFormatter {
	updateSubmateriFormatter := UpdateSubmateriFormatter{
		ID:   submateri.ID,
		Name: submateri.Name,
	}
	return updateSubmateriFormatter
}

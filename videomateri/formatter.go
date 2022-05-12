package videomateri

type CreateVideomateriFormatter struct {
	ID          int    `json:"id"`
	SubmateriID int    `json:"submateri_id"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

func FormatCreateVideomateri(videomateri VideomateriTable) CreateVideomateriFormatter {
	createVideomateriFormatter := CreateVideomateriFormatter{
		ID:          videomateri.ID,
		Name:        videomateri.Name,
		Link:        videomateri.Link,
		Description: videomateri.Description,
	}
	return createVideomateriFormatter
}

type UpdateVideomateriFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

func FormatUpdateVideomateri(videomateri VideomateriTable) UpdateVideomateriFormatter {
	updateVideomateriFormatter := UpdateVideomateriFormatter{
		ID:          videomateri.ID,
		Name:        videomateri.Name,
		Link:        videomateri.Link,
		Description: videomateri.Description,
	}
	return updateVideomateriFormatter
}

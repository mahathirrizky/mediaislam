package videomateri

type VideomateriFormatter struct {
	ID          int    `json:"id"`
	SubmateriID int    `json:"submateri_id"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

func FormatVideomateri(videomateri VideomateriTable) VideomateriFormatter {
	VideomateriFormatter := VideomateriFormatter{
		ID:          videomateri.ID,
		SubmateriID: videomateri.SubmateriID,
		Name:        videomateri.Name,
		Link:        videomateri.Link,
		Description: videomateri.Description,
	}
	return VideomateriFormatter
}

// type UpdateVideomateriFormatter struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Link        string `json:"link"`
// 	Description string `json:"description"`
// }

// func FormatUpdateVideomateri(videomateri VideomateriTable) UpdateVideomateriFormatter {
// 	updateVideomateriFormatter := UpdateVideomateriFormatter{
// 		ID:          videomateri.ID,
// 		Name:        videomateri.Name,
// 		Link:        videomateri.Link,
// 		Description: videomateri.Description,
// 	}
// 	return updateVideomateriFormatter
// }

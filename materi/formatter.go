package materi

type MateriFormatter struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	NamaUstadz  string `json:"nama_ustadz"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}

func FormatMateri(materi MateriTable) MateriFormatter {
	materiFormatter := MateriFormatter{
		ID:     materi.ID,
		UserID: materi.UserID,
		Name:        materi.Name,
		Description: materi.Description,
		Slug:        materi.Slug,
	}
	ustadz := materi.Ustadz
	materiFormatter.NamaUstadz = ustadz.Name

	return materiFormatter
}

func FormatMateriList(materi []MateriTable) []MateriFormatter {
	materilistFormatter := []MateriFormatter{}
	for _, materilist := range materi {
		materiFormatter := FormatMateri(materilist)
		materilistFormatter = append(materilistFormatter, materiFormatter)
	}
	return materilistFormatter
}

type SubmateriFormatter struct {
	ID          int    `json:"id"`
	Name 	  string `json:"name"`
}

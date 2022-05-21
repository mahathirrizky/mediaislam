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
		ID:          materi.ID,
		UserID:      materi.UserID,
		Name:        materi.Name,
		Description: materi.Description,
		Slug:        materi.Slug,
		NamaUstadz:  materi.Ustadz.Name,
	}

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

type MateriAllFormatter struct {
	ID        int `json:"id"`
	Name	  string `json:"name"`
	Submateri []SubmateriFormatter `json:"submateri"`
}

type SubmateriFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	VideoMateri []VideoMateriFormatter
}

type VideoMateriFormatter struct {
	ID          int    `json:"id"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

func FormatMateriAll(materi MateriTable) MateriAllFormatter {
	materiAllFormatter := MateriAllFormatter{
		ID: materi.ID,
		Name: materi.Name,
	}
	for _, submateri := range materi.Submateri {
		submateriFormatter := SubmateriFormatter{
			ID: submateri.ID,
			Name: submateri.Name,
		}
		for _, videomateri := range submateri.Videomateri {
			videoMateriFormatter := VideoMateriFormatter{
				ID:          videomateri.ID,
				Link:        videomateri.Link,
				Description: videomateri.Description,
			}
			submateriFormatter.VideoMateri = append(submateriFormatter.VideoMateri, videoMateriFormatter)
		}
		materiAllFormatter.Submateri = append(materiAllFormatter.Submateri, submateriFormatter)
	}
	return materiAllFormatter
}

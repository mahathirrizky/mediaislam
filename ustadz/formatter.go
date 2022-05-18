package ustadz

type ustadzFormatter struct {
	ID	int    `json:"id"`
	Name string `json:"name"`
}

func FormatUstadz(ustadz UstadzTable) ustadzFormatter {
	return ustadzFormatter{
		ID: ustadz.ID,
		Name: ustadz.Name,
	}
}

func FormatUstadzList(Ustadz []UstadzTable) []ustadzFormatter {
	ustadzlistFormatter := []ustadzFormatter{}
	for _, ustadzlist := range Ustadz {
		UstadzFormatter := FormatUstadz(ustadzlist)
		ustadzlistFormatter = append(ustadzlistFormatter, UstadzFormatter)
	}
	return ustadzlistFormatter
}
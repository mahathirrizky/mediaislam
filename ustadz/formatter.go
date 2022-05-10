package ustadz

type ustadzFormatter struct {
	Name string `json:"name"`
}

func FormatUstadz(ustadz UstadzTable) ustadzFormatter {
	return ustadzFormatter{
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
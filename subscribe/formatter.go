package subscribe

type SubscribeFormatter struct {
	ID         int    `json:"id"`
	MateriID   int    `json:"materi_id"`
	NamaMateri string `json:"nama_materi"`
}

func FormatSubscribe(subscribe SubscribeTable) SubscribeFormatter {
	subscribeFormatter := SubscribeFormatter{
		ID:       subscribe.ID,
		MateriID: subscribe.MateriID,
	}
	materi := subscribe.Materi
	subscribeFormatter.NamaMateri = materi.Name

	return subscribeFormatter
}

func FormatSubscribeList(subscribe []SubscribeTable) []SubscribeFormatter {
	SubscribelistFormatter := []SubscribeFormatter{}
	for _, subscribelist := range subscribe {
		SubscribeFormatter := FormatSubscribe(subscribelist)
		SubscribelistFormatter = append(SubscribelistFormatter, SubscribeFormatter)
	}
	return SubscribelistFormatter
}

type CreateSubscribeFormatter struct {
	ID       int `json:"id"`
	MateriID int `json:"materi_id"`
	UserID   int `json:"user_id"`
}

func FormatCreateSubscribe(subscribe SubscribeTable) CreateSubscribeFormatter {
	createSubscribeFormatter := CreateSubscribeFormatter{
		ID:       subscribe.ID,
		MateriID: subscribe.MateriID,
		UserID:   subscribe.UserID,
	}
	return createSubscribeFormatter
}

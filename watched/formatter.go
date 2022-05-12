package watched

type WatchedFormatter struct {
	ID            int `json:"id"`
	UserID        int `json:"user_id"`
	VideomateriID int `json:"videomateri_id"`
}

func FormatWatched(watched WatchedTable) WatchedFormatter {
	watchedFormatter := WatchedFormatter{
		ID:            watched.ID,
		UserID:        watched.UserID,
		VideomateriID: watched.VideomateriID,
	}

	return watchedFormatter
}

func FormatWatchedList(watched []WatchedTable) []WatchedFormatter {
	WatchedlistFormatter := []WatchedFormatter{}
	for _, watchedlist := range watched {
		watchedFormatter := FormatWatched(watchedlist)
		WatchedlistFormatter = append(WatchedlistFormatter, watchedFormatter)
	}
	return WatchedlistFormatter
}

type CreateWatchedFormatter struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	VideomateriID int `json:"videomateri_id"`
}

func FormatCreateWatched(watched WatchedTable) CreateWatchedFormatter {
	createWatchedFormatter := CreateWatchedFormatter{
		ID:       watched.ID,
		UserID:   watched.UserID,
		VideomateriID: watched.VideomateriID,
	}
	return createWatchedFormatter
}

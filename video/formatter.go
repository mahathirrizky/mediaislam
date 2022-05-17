package video

type VideoFormatter struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func FormatVideo(video VideoTable) VideoFormatter {
	VideoFormatter := VideoFormatter{
		ID:          video.ID,
		UserID:      video.UserID,
		Name:        video.Name,
		Link:        video.Link,
		Description: video.Description,
		ImageURL:    video.ImageFileName,
	}
	return VideoFormatter
}

func FormatVideoList(video []VideoTable) []VideoFormatter {
	VideoListFormatter := []VideoFormatter{}
	for _, videolist := range video {
		VideoFormatter := FormatVideo(videolist)
		VideoListFormatter = append(VideoListFormatter, VideoFormatter)
	}
	return VideoListFormatter
}


// type UpdateVideoFormatter struct {
// 	ID          int    `json:"id"`
// 	UserID      int    `json:"user_id"`
// 	Name        string `json:"name"`
// 	Link        string `json:"link"`
// 	Description string `json:"description"`
// }

// func FormatUpdateVideo(video VideoTable) UpdateVideoFormatter {
// 	updateVideoFormatter := UpdateVideoFormatter{
// 		ID:          video.ID,
// 		UserID:      video.UserID,
// 		Name:        video.Name,
// 		Link:        video.Link,
// 		Description: video.Description,
// 	}
// 	return updateVideoFormatter
// }

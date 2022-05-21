package video

import "fmt"

type service struct {
	repository Repository
}

type Service interface {
	CreateVideoTematik(input CreateVideoInput) (VideoTable, error)
	CreateVideoShort(input CreateVideoInput) (VideoTable, error)
	UpdateVideo(inputID GetVideoDetailInput, inputData CreateVideoInput) (VideoTable, error)
	GetVideo(inputID GetVideoDetailInput) (VideoTable, error)
	SaveImage(ID int, fileLocation string) (VideoTable, error)
	GetVideoTematik(userID int) ([]VideoTable, error)
	GetVideoShort(userID int) ([]VideoTable, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateVideoTematik(input CreateVideoInput) (VideoTable, error) {
	video := VideoTable{
		Name:        input.Name,
		Description: input.Description,
		Link:        input.Link,
		UserID:      input.User.ID,
		Type:		"tematik",
	}

	if input.User.Role != "contributor" {
		return video, fmt.Errorf("hanya contributor yang dapat membuat video")
	}

	newVideo, err := s.repository.Save(video)
	if err != nil {
		return newVideo, err
	}
	return newVideo, nil
}

func (s *service) CreateVideoShort(input CreateVideoInput) (VideoTable, error) {
	video := VideoTable{
		Name:        input.Name,
		Description: input.Description,
		Link:        input.Link,
		UserID:      input.User.ID,
		Type:		"short",
	}

	if input.User.Role != "contributor" {
		return video, fmt.Errorf("hanya contributor yang dapat membuat video")
	}

	newVideo, err := s.repository.Save(video)
	if err != nil {
		return newVideo, err
	}
	return newVideo, nil
}

func (s *service) UpdateVideo(inputID GetVideoDetailInput, inputData CreateVideoInput) (VideoTable, error) {
	video, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return video, err
	}
	if video.UserID != inputData.User.ID {
		return video, fmt.Errorf("user tidak bisa mengubah materi yang bukan miliknya")
	}
	video.Name = inputData.Name
	video.Link = inputData.Link
	video.Description = inputData.Description

	newVideo, err := s.repository.Update(video)
	if err != nil {
		return newVideo, err
	}
	return newVideo, nil
}

func (s *service) GetVideo(inputID GetVideoDetailInput) (VideoTable, error) {
	video, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return video, err
	}
	return video, nil
}

func (s *service) SaveImage(ID int, fileLocation string) (VideoTable, error) {
	video, err := s.repository.FindByID(ID)
	if err != nil {
		return video, err
	}
	video.ImageFileName = fileLocation
	imageVideo, err := s.repository.Update(video)
	if err != nil {
		return imageVideo, err
	}
	return imageVideo, nil
}

func (s *service) GetVideoTematik(userID int) ([]VideoTable, error) {
	if userID != 0 {
		tematik, err := s.repository.FindTematikByUserID(userID)
		if err != nil {
			return tematik, err
		}
		return tematik, nil
	}
	tematik, err := s.repository.FindTematik()
	if err != nil {
		return tematik, err
	}
	return tematik, nil
}

func (s *service) GetVideoShort(userID int) ([]VideoTable, error) {
	if userID != 0 {
		short, err := s.repository.FindShortByUserID(userID)
		if err != nil {
			return short, err
		}
		return short, nil
	}
	short, err := s.repository.FindShort()
	if err != nil {
		return short, err
	}
	return short, nil
}

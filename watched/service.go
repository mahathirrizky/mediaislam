package watched

type Service interface {
	GetWatched(userID int) ([]WatchedTable, error)
	CreateWatched(input CreateWatchedInput) (WatchedTable, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetWatched(userID int) ([]WatchedTable, error) {
	watcheds, err := s.repository.GetWatched(userID)
	if err != nil {
		return watcheds, err
	}
	return watcheds, nil
}

func (s *service) CreateWatched(input CreateWatchedInput) (WatchedTable, error) {
	watched := WatchedTable{
		VideomateriID: input.VideomateriID,
		UserID:        input.User.ID,
	}

	newWatched, err := s.repository.Save(watched)
	if err != nil {
		return newWatched, err
	}
	return newWatched, nil
}

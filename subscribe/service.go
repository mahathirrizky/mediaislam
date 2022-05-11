package subscribe

type Service interface {
	GetSubscribe(userID int) ([]SubscribeTable, error)
	CreateSubscribe(input CreateSubscribeInput) (SubscribeTable, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetSubscribe(userID int) ([]SubscribeTable, error) {
	subscribes, err := s.repository.GetSubscribe(userID)
	if err != nil {
		return subscribes, err
	}
	return subscribes, nil
}

func (s *service) CreateSubscribe(input CreateSubscribeInput) (SubscribeTable, error) {
	subscribe := SubscribeTable{
		MateriID: input.MateriID,
		UserID:   input.User.ID,
	}

	newSubscribe, err := s.repository.Save(subscribe)
	if err != nil {
		return newSubscribe, err
	}
	return newSubscribe, nil
}

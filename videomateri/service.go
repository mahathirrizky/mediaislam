package videomateri

type service struct {
	repository Repository
}

type Service interface {
	CreateVideomateri(input CreateVideomateriInput) (VideomateriTable, error)
	UpdateVideomateri(inputID GetVideomateriDetailInput, inputData CreateVideomateriInput) (VideomateriTable, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateVideomateri(input CreateVideomateriInput) (VideomateriTable, error) {
	videomateri := VideomateriTable{
		SubmateriID: input.SubmateriID,
		Link:        input.Link,
		Name:        input.Name,
		Description: input.Description,
	}

	newVideomateri, err := s.repository.Save(videomateri)
	if err != nil {
		return newVideomateri, err
	}
	return newVideomateri, nil
}

func (s *service) UpdateVideomateri(inputID GetVideomateriDetailInput, inputData CreateVideomateriInput) (VideomateriTable, error) {
	videomateri, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return videomateri, err
	}
	videomateri.Name = inputData.Name
	videomateri.Link = inputData.Link
	videomateri.Description = inputData.Description

	newVideomateri, err := s.repository.Update(videomateri)
	if err != nil {
		return newVideomateri, err
	}
	return newVideomateri, nil
}




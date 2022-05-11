package submateri

type service struct {
	repository Repository
}

type Service interface{
	CreateSubmateri(input CreateSubmateriInput) (SubmateriTable, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateSubmateri(input CreateSubmateriInput) (SubmateriTable, error) {
	submateri := SubmateriTable{
		MateriID: input.MateriID,
		Name: input.Name,
	}

	newSubmateri, err := s.repository.Save(submateri)
	if err != nil {
		return newSubmateri, err
	}
	return newSubmateri, nil
}
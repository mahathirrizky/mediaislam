package submateri

type service struct {
	repository Repository
}

type Service interface{
	CreateSubmateri(input CreateSubmateriInput) (SubmateriTable, error)
	UpdateSubmateri(inputID GetSubmateriDetailInput, inputData CreateSubmateriInput) (SubmateriTable, error)
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

func (s *service) UpdateSubmateri(inputID GetSubmateriDetailInput, inputData CreateSubmateriInput) (SubmateriTable, error) {
	submateri, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return submateri, err
	}
	submateri.Name = inputData.Name

	newSubmateri, err := s.repository.Update(submateri)
	if err != nil {
		return newSubmateri, err
	}
	return newSubmateri, nil
}
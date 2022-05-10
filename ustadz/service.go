package ustadz

type Service interface {
	RegisterUstadz(input RegisterUstadzInput) (UstadzTable, error)
	GetUstadz(ID int) ([]UstadzTable, error)
	GetUstadzByID(input GetUstadzDetailInput) (UstadzTable, error)
	UpdateUstadz(input GetUstadzDetailInput, inputData RegisterUstadzInput) (UstadzTable, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUstadz(input RegisterUstadzInput) (UstadzTable, error) {
	ustadz := UstadzTable{}
	ustadz.Name = input.Name

	newustadz, err := s.repository.Save(ustadz)
	if err != nil {
		return newustadz, err
	}
	return newustadz, nil

}

func (s *service) GetUstadzByID(input GetUstadzDetailInput) (UstadzTable, error) {
	ustadz, err := s.repository.FindByID(input.ID)
	if err != nil {
		return ustadz, err
	}
	return ustadz, nil
}

func (s *service) GetUstadz(ID int) ([]UstadzTable, error) {
	ustadz, err := s.repository.FindAll()
	if err != nil {
		return ustadz, err
	}
	return ustadz, nil
}

func (s *service) UpdateUstadz(input GetUstadzDetailInput, inputData RegisterUstadzInput) (UstadzTable, error) {
	ustadz, err := s.repository.FindByID(input.ID)

	if err != nil {
		return ustadz, err
	}

	ustadz.Name = inputData.Name

	updatedUstadz, err := s.repository.Update(ustadz)
	if err != nil {
		return updatedUstadz, err
	}
	return updatedUstadz, nil
}

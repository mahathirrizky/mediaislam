package materi

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetMateri(userID int) ([]MateriTable, error)
	GetMateriByID(input GetMateriDetailInput) (MateriTable, error)
	CreateMateri(input CreateMateriInput) (MateriTable, error)
	UpdateMateri(input GetMateriDetailInput, inputData CreateMateriInput) (MateriTable, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetMateri(userID int) ([]MateriTable, error) {
	if userID != 0 {
		materi, err := s.repository.FindByUserID(userID)
		if err != nil {
			return materi, err
		}
		return materi, nil
	}

	materi, err := s.repository.FindAll()
	if err != nil {
		return materi, err
	}
	return materi, nil

}

func (s *service) GetMateriByID(input GetMateriDetailInput) (MateriTable, error) {
	materi, err := s.repository.FindByID(input.ID)
	if err != nil {
		return materi, err
	}
	return materi, nil
}

func (s *service) CreateMateri(input CreateMateriInput) (MateriTable, error) {
	materi := MateriTable{
		UstadzID:    input.UstadzID,
		Name:        input.Name,
		Description: input.Description,
		UserID:      input.User.ID,
	}
	slugCandidate := fmt.Sprint("%s %d", input.Name, input.User.ID)
	materi.Slug = slug.Make(slugCandidate)

	newMateri, err := s.repository.Save(materi)
	if err != nil {
		return newMateri, err
	}
	return newMateri, nil
}

func (s *service) UpdateMateri(input GetMateriDetailInput, inputData CreateMateriInput) (MateriTable, error) {
	materi, err := s.repository.FindByID(input.ID)

	if err != nil {
		return materi, err
	}

	if materi.UserID != inputData.User.ID {
		return materi, fmt.Errorf("user tidak bisa mengubah materi yang bukan miliknya")
	}

	materi.UstadzID = inputData.UstadzID
	materi.Name = inputData.Name
	materi.Description = inputData.Description

	updatedMateri, err := s.repository.Update(materi)
	if err != nil {
		return updatedMateri, err
	}
	return updatedMateri, nil
}

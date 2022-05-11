package materi

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]MateriTable, error)
	FindByUserID(userID int) ([]MateriTable, error)
	Save(materi MateriTable) (MateriTable, error)
	FindByID(ID int) (MateriTable, error)
	FindMateriByUstadzID(ustadzID int) (MateriTable, error)
	Update(materi MateriTable) (MateriTable, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]MateriTable, error) {
	var materi []MateriTable
	err := r.db.Preload("Ustadz").Find(&materi).Error
	if err != nil {
		return materi, err
	}
	return materi, nil
}

func (r *repository) FindByUserID(userID int) ([]MateriTable, error) {
	var materi []MateriTable
	err := r.db.Where("user_id = ?", userID).Find(&materi).Error
	if err != nil {
		return materi, err
	}
	return materi, nil
}

func (r *repository) FindByID(ID int) (MateriTable, error) {
	var materi MateriTable
	err := r.db.Where("ID = ?", ID).Find(&materi).Error
	if err != nil {
		return materi, err
	}
	return materi, nil
}

func (r *repository) Save(materi MateriTable) (MateriTable, error) {
	err := r.db.Create(&materi).Error
	if err != nil {
		return materi, err
	}
	return materi, nil
}

func (r *repository) FindMateriByUstadzID(ustadzID int) (MateriTable, error) {
	var materi MateriTable
	err := r.db.Where("ustadz_id = ?", ustadzID).Find(&materi).Error
	if err != nil {
		return materi, err
	}
	return materi, nil
}

func (r *repository) Update(materi MateriTable) (MateriTable, error) {
	err := r.db.Save(&materi).Error
	if err != nil {
		return materi, err
	}
	return materi, nil
}

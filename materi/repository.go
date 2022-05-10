package materi

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]MateriTable, error)
	FindByUserID(userID int) ([]MateriTable, error)
	Save(materi MateriTable) (MateriTable, error)
	FindByID(ID int) (MateriTable, error)
	FindMateriByUstadzID(ustadzID int) (MateriTable, error)
	Update(materi MateriTable) (MateriTable, error)
	CreateSubmateri(submateri SubmateriTable) (SubmateriTable, error)
	FindByMateriID(materiID int) (SubmateriTable, error)
	FindBySubmateriID(submateriID int) (VideomateriTable, error)
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

func (r *repository) CreateSubmateri(submateri SubmateriTable) (SubmateriTable, error) {
	err := r.db.Create(&submateri).Error
	if err != nil {
		return submateri, err
	}
	return submateri, nil
}

func (r *repository) FindByMateriID(materiID int) (SubmateriTable, error) {
	var submateri SubmateriTable
	err := r.db.Where("materi_id = ?", materiID).Find(&submateri).Error
	if err != nil {
		return submateri, err
	}
	return submateri, nil
}

func (r *repository) FindBySubmateriID(submateriID int) (VideomateriTable, error) {
	var videomateri VideomateriTable
	err := r.db.Where("submateri_id = ?", submateriID).Find(&videomateri).Error
	if err != nil {
		return videomateri, err
	}
	return videomateri, nil
}

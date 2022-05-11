package submateri

import "gorm.io/gorm"

type Repository interface {
	Save(submateri SubmateriTable) (SubmateriTable, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(submateri SubmateriTable) (SubmateriTable, error){
	err := r.db.Create(&submateri).Error
	if err != nil {
		return submateri, err
	}
	return submateri, nil
}
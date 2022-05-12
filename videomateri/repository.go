package videomateri

import "gorm.io/gorm"

type Repository interface {
	Save(videomateri VideomateriTable) (VideomateriTable, error)
	Update(videomateri VideomateriTable) (VideomateriTable, error)
	FindByID(ID int) (VideomateriTable, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(videomateri VideomateriTable) (VideomateriTable, error) {
	err := r.db.Create(&videomateri).Error
	if err != nil {
		return videomateri, err
	}
	return videomateri, nil
}

func (r *repository) Update(videomateri VideomateriTable) (VideomateriTable, error) {
	err := r.db.Save(&videomateri).Error
	if err != nil {
		return videomateri, err
	}
	return videomateri, nil
}

func (r *repository) FindByID(ID int) (VideomateriTable, error) {
	var videomateri VideomateriTable
	err := r.db.Where("id = ?", ID).First(&videomateri).Error
	if err != nil {
		return videomateri, err
	}
	return videomateri, nil
}

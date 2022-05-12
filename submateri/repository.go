package submateri

import "gorm.io/gorm"

type Repository interface {
	Save(submateri SubmateriTable) (SubmateriTable, error)
	Update(submateri SubmateriTable) (SubmateriTable, error)
	FindByID(ID int) (SubmateriTable, error)
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

 func (r *repository) Update(submateri SubmateriTable) (SubmateriTable, error){
	err := r.db.Save(&submateri).Error
	if err != nil {
		return submateri, err
	}	
	return submateri, nil
}

 func (r *repository) FindByID(ID int) (SubmateriTable, error){
	var submateri SubmateriTable
	err := r.db.Where("id = ?", ID).First(&submateri).Error
	if err != nil {
		return submateri, err
	}
	return submateri, nil
 }
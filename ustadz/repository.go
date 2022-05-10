package ustadz

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]UstadzTable, error)
	Save(ustadz UstadzTable) (UstadzTable, error)
	FindByID(ID int) (UstadzTable, error)
	Update(ustadz UstadzTable) (UstadzTable, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(ustadz UstadzTable) (UstadzTable, error) {
	err := r.db.Create(&ustadz).Error
	if err != nil {
		return ustadz, err
	}
	return ustadz, nil
}

func (r *repository) FindAll() ([]UstadzTable, error) {
	var ustadz []UstadzTable
	err := r.db.Find(&ustadz).Error
	if err != nil {
		return ustadz, err
	}
	return ustadz, nil
}

func (r *repository) FindByEmail(email string) (UstadzTable, error){
	var ustadz UstadzTable

	err := r.db.Where("email = ?", email).Find(&ustadz).Error
	if err != nil{
		return ustadz, err
	}
	return ustadz, nil
}

func (r *repository) FindByID(ID int) (UstadzTable, error){
	var ustadz UstadzTable

	err := r.db.Where("ID = ?", ID).Find(&ustadz).Error
	if err != nil{
		return ustadz, err
	}
	return ustadz, nil
}

func (r *repository) Update(ustadz UstadzTable) (UstadzTable, error){
	err := r.db.Save(&ustadz).Error
	if err != nil{
		return ustadz, err
	}
	return ustadz, nil
}

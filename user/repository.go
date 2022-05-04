package user

import "gorm.io/gorm"

type Repository interface {
	Save(user UserTable) (UserTable, error)
	FindByEmail(email string) (UserTable, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user UserTable) (UserTable, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (UserTable, error){
	var user UserTable

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil{
		return user, err
	}
	return user, nil
}

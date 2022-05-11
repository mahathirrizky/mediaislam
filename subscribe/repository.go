package subscribe

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetSubscribe(useriID int) ([]SubscribeTable, error)
	Save(subscribe SubscribeTable) (SubscribeTable, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetSubscribe(useriID int) ([]SubscribeTable, error) {
	var subscribes []SubscribeTable
	err := r.db.Preload("Materi").Where("user_id = ?", useriID).Find(&subscribes).Error
	if err != nil {
		return subscribes, err
	}
	return subscribes, nil
}

func (r *repository) Save(subscribe SubscribeTable) (SubscribeTable, error) {
	err := r.db.Create(&subscribe).Error
	if err != nil {
		return subscribe, err
	}
	return subscribe, nil
}

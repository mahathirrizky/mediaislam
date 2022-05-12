package watched

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetWatched(userID int) ([]WatchedTable, error)
	Save(watched WatchedTable) (WatchedTable, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetWatched(userID int) ([]WatchedTable, error) {
	var watcheds []WatchedTable
	err := r.db.Preload("Videomateri").Where("user_id = ?", userID).Find(&watcheds).Error
	if err != nil {
		return watcheds, err
	}
	return watcheds, nil
}

func (r *repository) Save(watched WatchedTable) (WatchedTable, error) {
	err := r.db.Create(&watched).Error
	if err != nil {
		return watched, err
	}
	return watched, nil
}
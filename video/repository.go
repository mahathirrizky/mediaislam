package video

import "gorm.io/gorm"

type Repository interface {
	Save(video VideoTable) (VideoTable, error)
	Update(video VideoTable) (VideoTable, error)
	FindByID(ID int) (VideoTable, error)
	FindTematik() ([]VideoTable, error)
	FindTematikByUserID(userID int) ([]VideoTable, error)
	FindShort() ([]VideoTable, error)
	FindShortByUserID(userID int) ([]VideoTable, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(video VideoTable) (VideoTable, error) {
	err := r.db.Create(&video).Error
	if err != nil {
		return video, err
	}
	return video, nil
}

func (r *repository) Update(video VideoTable) (VideoTable, error) {
	err := r.db.Save(&video).Error
	if err != nil {
		return video, err
	}
	return video, nil
}

func (r *repository) FindByID(ID int) (VideoTable, error) {
	var video VideoTable
	err := r.db.Where("id = ?", ID).First(&video).Error
	if err != nil {
		return video, err
	}
	return video, nil
}

func (r *repository) FindTematik() ([]VideoTable, error) {
	var videos []VideoTable
	err := r.db.Where("type = ?", "tematik").Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

func (r *repository) FindTematikByUserID(userID int) ([]VideoTable, error) {
	var video []VideoTable
	err := r.db.Where("user_id = ? AND type = ?", userID, "tematik").Find(&video).Error
	if err != nil {
		return video, err
	}
	return video, nil
}

func (r *repository) FindShort() ([]VideoTable, error) {
	var videos []VideoTable
	err := r.db.Where("type = ?", "short").Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

func (r *repository) FindShortByUserID(userID int) ([]VideoTable, error) {
	var video []VideoTable
	err := r.db.Where("user_id = ? AND type = ?", userID, "short").Find(&video).Error
	if err != nil {
		return video, err
	}
	return video, nil
}

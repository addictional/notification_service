package notification

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(notification *Notification) error {
	return r.db.Create(notification).Error
}

func (r *Repository) List() ([]Notification, error) {
	var notifications []Notification
	if err := r.db.Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

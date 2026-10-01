package notification

import "time"

type Notification struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	UserID    int64     `json:"user_id" gorm:"not null"`
	Message   string    `json:"message" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateNotificationRequest struct {
	UserID  int64  `json:"user_id"`
	Message string `json:"message"`
}

type ListNotificationsDto struct {
	Notifications []Notification `json:"items"`
	Total         int64          `json:"total"`
}

package repository

import (
	"realtime/internals/Config"
	"realtime/internals/Models"
)

type NotificationRepository struct{}

func NewNotificationRepository() *NotificationRepository {
	return &NotificationRepository{}
}

func (r *NotificationRepository) Create(notification *models.Notification) error {
	return config.DB.Create(notification).Error
}

func (r *NotificationRepository) GetForUser(userID uint64) ([]models.Notification, error) {
	var notifications []models.Notification
	err := config.DB.Where("notifiable_id = ? AND notifiable_type = ?", userID, "Admin").Find(&notifications).Error
	return notifications, err
}

func (r *NotificationRepository) MarkAsRead(id string) error {
	return config.DB.Model(&models.Notification{}).Where("id = ?", id).Update("read_at", "NOW()").Error
}

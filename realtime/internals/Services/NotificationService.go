package services

import (
	"realtime/internals/Models"
	"realtime/internals/Repository"
)

type NotificationService struct {
	repo *repository.NotificationRepository
}

func NewNotificationService(repo *repository.NotificationRepository) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) SendNotification(userID uint64, nType string, data string) (*models.Notification, error) {
	notification := &models.Notification{
		Type:           nType,
		NotifiableType: "Admin",
		NotifiableID:   userID,
		Data:           data,
	}

	err := s.repo.Create(notification)
	return notification, err
}

func (s *NotificationService) GetUserNotifications(userID uint64) ([]models.Notification, error) {
	return s.repo.GetForUser(userID)
}

package websockets

import (
	"realtime/internals/commen"
)

type Notification struct {
	ID      int32   `json:"id"`
	Message string  `json:"message"`
	Targets []int32 `json:"targets"`
}

type NotificationService struct {
	clients    map[int32]*commen.Client
	Register   chan *commen.Client
	Unregister chan *commen.Client
	Broadcast  chan *Notification
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		clients:    make(map[int32]*commen.Client),
		Register:   make(chan *commen.Client),
		Unregister: make(chan *commen.Client),
		Broadcast:  make(chan *Notification),
	}
}
func (s *NotificationService) Run() {
	for {
		select {
		case client := <-s.Register:
			s.clients[client.ID] = client
			go client.WritePump()
			client.Inbox <- []byte("Welcome to the notification service!")
		case client := <-s.Unregister:
			delete(s.clients, client.ID)
		case notification := <-s.Broadcast:
			for _, target := range notification.Targets {
				if client, ok := s.clients[target]; ok {
					client.Inbox <- []byte(notification.Message)
				}
			}
		}
	}
}
func (s *NotificationService) SendNotification(notification *Notification) {
	s.Broadcast <- notification
}
func (s *NotificationService) RegisterClient(client *commen.Client) {
	s.Register <- client
}
func (s *NotificationService) UnregisterClient(client *commen.Client) {
	s.Unregister <- client
}

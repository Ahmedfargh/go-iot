package models

import (
	Notification "go-dashboard/internals/Gprc/notification"
	"log"

	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name        string       `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Email       string       `gorm:"column:email;type:varchar(32);not null;unique" json:"email"`
	Phone       string       `gorm:"column:phone;type:varchar(15);not null;unique" json:"phone"`
	Password    string       `gorm:"column:password;type:varchar(150);not null" json:"password"`
	Avatar      string       `gorm:"column:avatar;type:varchar(100)" json:"avatar"`
	Roles       []Role       `gorm:"many2many:admin_roles;" json:"roles"`
	Permissions []Permission `gorm:"many2many:admin_permissions;" json:"permissions"`
}

var ns_admin_notification Notification.NotificationServiceClient

func InitAuthClient() {
	creds, err := credentials.NewClientTLSFromFile("server.crt", "localhost")
	if err != nil {
		log.Fatalf("Failed to load TLS certs: %v", err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}
	ns_admin_notification = Notification.NewNotificationServiceClient(conn)
}
func (admin *Admin) AfterCreate(tx *gorm.DB) error {
	var adminIDs []uint

	err := tx.Table("admins").
		Select("admins.id").
		Joins("JOIN admin_roles ON admin_roles.admin_id = admins.id").
		Joins("JOIN roles ON roles.id = admin_roles.role_id").
		Where("roles.name = ?", "admin").
		Pluck("admins.id", &adminIDs).Error
	ctx := context.Background()
	// convert []uint -> []int32 for proto
	ids := make([]int32, len(adminIDs))
	for i, v := range adminIDs {
		ids[i] = int32(v)
	}

	message_param := Notification.NotificationParam{
		Message: "new admin called " + admin.Name + " has been created welcome him",
		Ids:     ids,
	}
	ns_admin_notification.Notifiy(ctx, &message_param)
	return err
}

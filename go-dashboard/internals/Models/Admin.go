package models

import (
	Notification "go-dashboard/internals/Gprc/notification"
	"log"

	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gorm.io/gorm"
	"fmt"
)

type Admin struct {
	gorm.Model
	Name        string       `gorm:"column:name;type:varchar(50);not null" json:"name" form:"name"`
	Email       string       `gorm:"column:email;type:varchar(32);not null;unique" json:"email" form:"email"`
	Phone       string       `gorm:"column:phone;type:varchar(15);not null;unique" json:"phone" form:"phone"`
	Password    string       `gorm:"column:password;type:varchar(150);not null" json:"password" form:"password"`
	Avatar      string       `gorm:"column:avatar;type:varchar(100)" json:"avatar" form:"avatar"`
	Roles       []Role       `gorm:"many2many:admin_roles;" json:"roles" form:"roles"`
	Permissions []Permission `gorm:"many2many:admin_permissions;" json:"permissions" form:"permissions"`
}


func InitNotificationClient() Notification.NotificationServiceClient{
	var ns_admin_notification Notification.NotificationServiceClient

	creds, err := credentials.NewClientTLSFromFile("server.crt", "localhost")
	if err != nil {
		log.Fatalf("Failed to load TLS certs: %v", err)
	}

	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}
	ns_admin_notification = Notification.NewNotificationServiceClient(conn)
	return ns_admin_notification
}
func (admin *Admin) AfterCreate(tx *gorm.DB) error {
	var adminIDs []uint
	fmt.Println("GORM HOOK CALLED !!!!!!!!")
	err := tx.Table("admins").
		Select("admins.id").
		Joins("JOIN admin_roles ON admin_roles.admin_id = admins.id").
		Joins("JOIN roles ON roles.id = admin_roles.role_id").
		Where("roles.name = ?", "SuperAdmin").
		Pluck("admins.id", &adminIDs).Error
	if err!=nil{
		fmt.Println("ERROR GETTING ADMIN")
	}
	ctx := context.Background()
	// convert []uint -> []int32 for proto
	ids := make([]int32, len(adminIDs))
	for i, v := range adminIDs {
		ids[i] = int32(v)
	}
	fmt.Println("print admin ids:-",ids)
	message_param := Notification.NotificationParam{
		Message: "new admin called " + admin.Name + " has been created welcome him",
		Ids:     ids,
	}
	ns_admin_notification:=InitNotificationClient()
	fmt.Println("fmt prinln before notify")
	ns_admin_notification.Notifiy(ctx, &message_param)
	return err
}

package Gprc

import (
	"context"
	"fmt"
	"log"
	"net"
	Notification "realtime/internals/Gprc/notification"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"realtime/internals/WebSockets"
)

type NotificationService struct {
	Notification.UnimplementedNotificationServiceServer
}

func (ns *NotificationService) Notifiy(ctx context.Context, param *Notification.NotificationParam) (*Notification.NotificationReturn, error) {
    fmt.Println("notificationReceived")
    notification := websockets.Notification{
        Message: param.Message, 
        Targets: param.Ids,
    }
    websockets.Notification_service.SendNotification(&notification)
    fmt.Println("Notification Served")
    return &Notification.NotificationReturn{
		RecievedIds:make([]int32, 5),
		NotRecievedIds:make([]int32, 5),
	}, nil
}
func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 1. Load your TLS credentials
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		fmt.Println("test")
		log.Fatalf("failed to load certificates: %v", err)
	}

	// 2. Pass the creds to the gRPC server options
	s := grpc.NewServer(grpc.Creds(creds))
	notification_service:=NotificationService{}
	Notification.RegisterNotificationServiceServer(s, &notification_service)

	fmt.Println("gRPC Auth Server listening on :50052 (Secure)")
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
		log.Fatalf("failed to serve: %v", err)
	}
}

package Gprc

import (
	"context"
	"fmt"
	"log"
	"net"
	Notification "realtime/internals/Gprc/notification"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type NotificationService struct {
	Notification.UnimplementedNotificationServiceServer
}

func (ns *NotificationService) Notifiy(ctx context.Context, notification_param *Notification.NotificationParam) (*Notification.NotificationReturn, error) {
	fmt.Println("notificationReceived")
	return nil, nil
}
func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 1. Load your TLS credentials
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatalf("failed to load certificates: %v", err)
	}

	// 2. Pass the creds to the gRPC server options
	s := grpc.NewServer(grpc.Creds(creds))

	Notification.RegisterNotificationServiceServer(s, &NotificationService{})

	fmt.Println("gRPC Auth Server listening on :50051 (Secure)")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

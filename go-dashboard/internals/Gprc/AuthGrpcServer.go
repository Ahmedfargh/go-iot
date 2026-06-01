package gprc

import (
	"context"
	"fmt"
	"log"
	"net"

	Autoloader "go-dashboard/internals/AutoLoader"
	"go-dashboard/internals/Gprc/authpb"
	authServices "go-dashboard/internals/Services/Auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
}

func (s *AuthServer) IsAuth(ctx context.Context, req *authpb.AuthRequest) (*authpb.AuthResponse, error) {
	authService := Autoloader.Services_providers["admin.auth"].(*authServices.AdminAuthService)
	admin, err := authService.ValidateToken(req.Token)
	if err != nil {
		return &authpb.AuthResponse{
			Authenticated: false,
			Error:         err.Error(),
		}, nil
	}

	return &authpb.AuthResponse{
		Authenticated: true,
		UserId:        uint64(admin.ID),
	}, nil
}

func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":50051")
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

	authpb.RegisterAuthServiceServer(s, &AuthServer{})

	fmt.Println("gRPC Auth Server listening on :50051 (Secure)")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

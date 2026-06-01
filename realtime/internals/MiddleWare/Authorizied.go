package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	authpb "realtime/internals/Gprc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Global client to reuse the connection
var authClient authpb.AuthServiceClient

func InitAuthClient() {
	creds, err := credentials.NewClientTLSFromFile("server.crt", "localhost")
	if err != nil {
		log.Fatalf("Failed to load TLS certs: %v", err)
	}

	// Use grpc.Dial or NewClient depending on your grpc version
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}
	authClient = authpb.NewAuthServiceClient(conn)
}

func Authorizied() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := authClient.IsAuth(ctx, &authpb.AuthRequest{Token: parts[1]})
		if err != nil || !resp.Authenticated {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user_id", resp.UserId)
		c.Next()
	}
}

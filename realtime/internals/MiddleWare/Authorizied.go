package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"realtime/internals/Gprc/authpb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Authorizied() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must be Bearer token"})
			c.Abort()
			return
		}

		token := parts[1]

		// Connect to go-dashboard gRPC server
		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to auth service"})
			c.Abort()
			return
		}
		defer conn.Close()

		client := authpb.NewAuthServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := client.IsAuth(ctx, &authpb.AuthRequest{Token: token})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Auth service unavailable"})
			c.Abort()
			return
		}

		if !resp.Authenticated {
			c.JSON(http.StatusUnauthorized, gin.H{"error": resp.Error})
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", resp.UserId)
		c.Next()
	}
}

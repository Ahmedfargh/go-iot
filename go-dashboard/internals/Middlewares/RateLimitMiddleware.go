package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type client struct {
	lastSeen time.Time
	count    int
}

var (
	clients = make(map[string]*client)
	mu      sync.Mutex
)

// RateLimitMiddleware limits requests per IP
// limit: max requests, window: time window for the limit
func RateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {
	// Cleanup routine
	go func() {
		for {
			time.Sleep(window)
			mu.Lock()
			for ip, c := range clients {
				if time.Since(c.lastSeen) > window {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		defer mu.Unlock()

		if _, found := clients[ip]; !found {
			clients[ip] = &client{lastSeen: time.Now(), count: 1}
			c.Next()
			return
		}

		if time.Since(clients[ip].lastSeen) > window {
			clients[ip].count = 1
			clients[ip].lastSeen = time.Now()
			c.Next()
			return
		}

		if clients[ip].count >= limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		clients[ip].count++
		c.Next()
	}
}

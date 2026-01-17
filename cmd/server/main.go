package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/slack-clone-backend/internal/config"
	"github.com/yourusername/slack-clone-backend/internal/database"
	"github.com/yourusername/slack-clone-backend/pkg/jwt"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Set Gin mode
	gin.SetMode(cfg.GinMode)

	// Connect to PostgreSQL
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Connect to Redis
	redis, err := database.NewRedisClient(cfg.RedisURL, cfg.RedisPassword)
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	defer redis.Close()

	// Initialize JWT manager
	jwtManager := jwt.NewJWTManager(
		cfg.JWTSecret,
		cfg.JWTAccessExpiry,
		cfg.JWTRefreshExpiry,
	)

	// Create Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	// API routes
	api := router.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", func(c *gin.Context) {
				// TODO: Implement register handler
				c.JSON(http.StatusOK, gin.H{"message": "Register endpoint - TODO"})
			})
			auth.POST("/login", func(c *gin.Context) {
				// TODO: Implement login handler
				c.JSON(http.StatusOK, gin.H{"message": "Login endpoint - TODO"})
			})
			auth.POST("/refresh", func(c *gin.Context) {
				// TODO: Implement refresh handler
				c.JSON(http.StatusOK, gin.H{"message": "Refresh endpoint - TODO"})
			})
			auth.POST("/logout", func(c *gin.Context) {
				// TODO: Implement logout handler
				c.JSON(http.StatusOK, gin.H{"message": "Logout endpoint - TODO"})
			})
		}

		// Protected routes (require authentication)
		protected := api.Group("")
		// protected.Use(authMiddleware(jwtManager))
		{
			// User routes
			users := protected.Group("/users")
			{
				users.GET("/me", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"message": "Get current user - TODO"})
				})
				users.PUT("/me", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"message": "Update user - TODO"})
				})
			}

			// Workspace routes
			workspaces := protected.Group("/workspaces")
			{
				workspaces.GET("", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"message": "List workspaces - TODO"})
				})
				workspaces.POST("", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"message": "Create workspace - TODO"})
				})
				workspaces.GET("/:id", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"message": "Get workspace - TODO"})
				})
			}

			// Channel routes
			channels := protected.Group("/channels")
			{
				channels.GET("/:id/messages", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"message": "Get messages - TODO"})
				})
				channels.POST("/:id/messages", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"message": "Send message - TODO"})
				})
			}
		}
	}

	// WebSocket endpoint
	router.GET("/ws", func(c *gin.Context) {
		// TODO: Implement WebSocket handler
		c.JSON(http.StatusOK, gin.H{"message": "WebSocket endpoint - TODO"})
	})

	// WebRTC signaling endpoint
	router.GET("/webrtc/signaling", func(c *gin.Context) {
		// TODO: Implement WebRTC signaling handler
		c.JSON(http.StatusOK, gin.H{"message": "WebRTC signaling - TODO"})
	})

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("🚀 Server starting on %s", addr)
	log.Printf("📝 Gin mode: %s", cfg.GinMode)
	log.Printf("🌐 Allowed origins: %v", cfg.AllowedOrigins)

	if err := router.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

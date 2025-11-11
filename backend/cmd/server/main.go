package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shafikshaon/url_shortener/config"
	"github.com/shafikshaon/url_shortener/internal/analytics"
	"github.com/shafikshaon/url_shortener/internal/api"
	"github.com/shafikshaon/url_shortener/internal/auth"
	"github.com/shafikshaon/url_shortener/internal/database"
	"github.com/shafikshaon/url_shortener/internal/service"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("✓ Database connected successfully")

	// Initialize repositories
	userRepo := database.NewUserRepository(db.DB)
	linkRepo := database.NewLinkRepository(db.DB)
	analyticsRepo := database.NewAnalyticsRepository(db.DB)

	// Initialize services
	jwtService := auth.NewJWTService(cfg)
	linkService := service.NewLinkService(linkRepo, userRepo)
	tracker := analytics.NewTracker(analyticsRepo)

	// Initialize handlers
	authHandler := api.NewAuthHandler(userRepo, jwtService)
	linkHandler := api.NewLinkHandler(linkService, tracker, cfg)
	analyticsHandler := api.NewAnalyticsHandler(tracker)

	// Setup Gin router
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// CORS configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	// Public routes
	router.GET("/:code", linkHandler.Redirect)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Auth routes (public)
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/signup", authHandler.Signup)
			authRoutes.POST("/login", authHandler.Login)
		}

		// Protected routes (require JWT)
		protected := v1.Group("")
		protected.Use(auth.AuthMiddleware(jwtService))
		{
			// User routes
			protected.GET("/profile", authHandler.GetProfile)
			protected.POST("/api-key", authHandler.GenerateAPIKey)

			// Link routes
			protected.POST("/links", linkHandler.CreateLink)
			protected.GET("/links", linkHandler.ListLinks)
			protected.GET("/links/:id", linkHandler.GetLink)
			protected.PATCH("/links/:id", linkHandler.UpdateLink)
			protected.DELETE("/links/:id", linkHandler.DeleteLink)
			protected.GET("/links/:id/stats", linkHandler.GetLinkStats)

			// Tag routes
			protected.GET("/tags", linkHandler.GetUserTags)

			// Analytics routes
			protected.GET("/analytics", analyticsHandler.GetUserAnalytics)
		}

		// API Key protected routes (for external API access)
		apiKeyProtected := v1.Group("/api")
		apiKeyProtected.Use(auth.APIKeyMiddleware(userRepo))
		{
			apiKeyProtected.POST("/links", linkHandler.CreateLink)
			apiKeyProtected.GET("/links", linkHandler.ListLinks)
		}
	}

	// Start server
	addr := ":" + cfg.Server.Port
	log.Printf("🚀 Server starting on %s", addr)
	log.Printf("📊 Environment: %s", cfg.Env)
	log.Printf("🔗 Base URL: %s", cfg.Server.BaseURL)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

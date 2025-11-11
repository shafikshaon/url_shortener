package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shafikshaon/url_shortener/config"
	"github.com/shafikshaon/url_shortener/internal/analytics"
	"github.com/shafikshaon/url_shortener/internal/api"
	"github.com/shafikshaon/url_shortener/internal/auth"
	"github.com/shafikshaon/url_shortener/internal/database"
	"github.com/shafikshaon/url_shortener/internal/logger"
	"github.com/shafikshaon/url_shortener/internal/middleware"
	"github.com/shafikshaon/url_shortener/internal/service"
)

func main() {
	ctx := context.Background()

	// Load configuration
	cfg := config.Load()
	logger.Infof(ctx, "Configuration loaded for environment: %s", cfg.Env)

	// Connect to GORM database
	gormDB, err := database.NewGormDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer gormDB.Close()

	logger.Infof(ctx, "✓ Database connected successfully")

	// Run auto-migrations
	if err := gormDB.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	logger.Infof(ctx, "✓ Database migrations completed")

	// Initialize GORM repositories
	userRepo := database.NewUserRepository(gormDB.DB)
	linkRepo := database.NewLinkRepository(gormDB.DB)
	analyticsRepo := database.NewAnalyticsRepository(gormDB.DB)

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

	// Add trace middleware to inject trace IDs
	router.Use(middleware.TraceMiddleware())
	logger.Infof(ctx, "✓ Trace middleware enabled")

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
	logger.Infof(ctx, "🚀 Server starting on %s", addr)
	logger.Infof(ctx, "📊 Environment: %s", cfg.Env)
	logger.Infof(ctx, "🔗 Base URL: %s", cfg.Server.BaseURL)

	if err := router.Run(addr); err != nil {
		logger.Errorf(ctx, "Failed to start server: %+v", err)
		log.Fatalf("Failed to start server: %v", err)
	}
}

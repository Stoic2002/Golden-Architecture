package main

import (
	"log"

	"github.com/arulkarim/golden-architecture/configs"
	_ "github.com/arulkarim/golden-architecture/docs" // Swagger docs
	"github.com/arulkarim/golden-architecture/internal/infrastructure/auth"
	"github.com/arulkarim/golden-architecture/internal/infrastructure/database"
	infrahttp "github.com/arulkarim/golden-architecture/internal/infrastructure/http"
	"github.com/arulkarim/golden-architecture/internal/todo"
	todohandler "github.com/arulkarim/golden-architecture/internal/todo/handler"
	todopostgres "github.com/arulkarim/golden-architecture/internal/todo/postgres"
	"github.com/arulkarim/golden-architecture/internal/user"
	userhandler "github.com/arulkarim/golden-architecture/internal/user/handler"
	userpostgres "github.com/arulkarim/golden-architecture/internal/user/postgres"
	"github.com/arulkarim/golden-architecture/pkg/validator"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load configuration
	cfg, err := configs.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := database.NewPostgresDB(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run auto migration
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migration: %v", err)
	}

	// Initialize validator
	validator.Init()

	// Initialize JWT manager
	jwtManager := auth.NewJWTManager(&cfg.JWT)

	// Wire Todo dependencies
	todoRepo := todopostgres.NewTodoRepository(db)
	todoService := todo.NewService(todoRepo)
	todoHandler := todohandler.NewHandler(todoService)

	// Wire User/Auth dependencies
	userRepo := userpostgres.NewUserRepository(db)
	userService := user.NewService(userRepo, jwtManager)
	userHandler := userhandler.NewHandler(userService)

	// Create HTTP server
	server := infrahttp.NewServer(cfg.Server.Port, cfg.Server.Mode)

	// Register routes
	api := server.Engine().Group("/api/v1")
	todohandler.RegisterRoutes(api, todoHandler)
	userhandler.RegisterRoutes(api, userHandler, jwtManager)

	// Swagger documentation endpoint
	server.Engine().GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint
	server.Engine().GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// Start server
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

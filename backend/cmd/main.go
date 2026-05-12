package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/industrix-todo/backend/internal/handler"
	"github.com/industrix-todo/backend/internal/repository"
	"github.com/industrix-todo/backend/internal/service"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL must be set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	categoryRepo := repository.NewCategoryRepository(db)
	todoRepo := repository.NewTodoRepository(db)

	// Initialize services
	categoryService := service.NewCategoryService(categoryRepo)
	todoService := service.NewTodoService(todoRepo)

	// Initialize handlers
	categoryHandler := handler.NewCategoryHandler(categoryService)
	todoHandler := handler.NewTodoHandler(todoService)

	// Setup Gin
	r := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // For development
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Routes
	api := r.Group("/api")
	categoryHandler.RegisterRoutes(api)
	todoHandler.RegisterRoutes(api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\cmd\server\main.go






// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\cmd\server\main.go





package main

import (
    "log"
    "Backend/config"               // Local import for config
    "Backend/internal/api/router"  // Local import for router
    "Backend/internal/service"     // Local import for service layer
    "Backend/internal/store"       // Local import for store interface
    "Backend/internal/store/memory" // Local import for memory store
    "Backend/internal/store/postgres" // Local import for postgres store
    "Backend/internal/store/redis" // Local import for redis store
    "github.com/gin-gonic/gin"     // Added Gin import
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title URL Shortener API
// @version 1.0
// @description A simple URL shortening service built with Go and Gin.
// @host localhost:8080
// @BasePath /
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
func main() {
    // Initialize Gin router
    r := gin.Default()

    // Load the configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    var urlService *service.URLService
    var store store.Store // The store should be of type service.Store

    // Choose the store based on the configuration
    switch cfg.StoreType {
    case "memory":
        // Initialize the MemoryStore with file persistence
        store = memory.New() // Initialize the MemoryStore without arguments
    case "postgres":
        store, err = postgres.New(cfg.PostgresDSN) // PostgreSQL connection string from config
        if err != nil {
            log.Fatalf("Failed to connect to PostgreSQL: %v", err)
        }
    case "redis":
        store, err = redis.New(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB) // Redis connection info from config
        if err != nil {
            log.Fatalf("Failed to connect to Redis: %v", err)
        }
    default:
        log.Fatalf("Invalid store type: %v", cfg.StoreType)
    }

    // Initialize the service with the selected store
    urlService = service.NewURLService(store)

    // Setup the router with the URL service, passing the existing router
    router.SetupRouter(r, urlService) // Fixed to pass r, not reassign

    // Add health check endpoint
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy"})
    })

    // Add Swagger endpoint with custom configuration
    config := &ginSwagger.Config{
        URL: "http://localhost:8080/docs/swagger.json", // Explicitly point to swagger.json
    }
    r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

    // Serve the swagger.json file directly
    r.StaticFile("/docs/swagger.json", "./docs/swagger.json")

    // Start the server on the configured port
    if err := r.Run(":" + cfg.ServerPort); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
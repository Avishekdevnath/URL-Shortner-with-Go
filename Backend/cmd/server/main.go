// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\cmd\server\main.go
package main

import (
	"time"
	"Backend/config"
	"Backend/internal/api/router"
	"Backend/internal/service"
	"Backend/internal/store"
	"Backend/internal/store/memory"
	"Backend/internal/store/postgres"
	"Backend/internal/store/redis"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"  // Import ginSwagger for Swagger UI
    "github.com/swaggo/files" 
)

var (
	requestLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Latency of HTTP requests",
		Buckets: prometheus.LinearBuckets(0.01, 0.05, 10),
	}, []string{"path"})
)

func init() {
	prometheus.MustRegister(requestLatency)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	r := gin.Default()

	logrus.Info("Starting URL Shortener server...")

	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("Error loading config: %v", err)
	}

	var storeInstance store.Store
	var redisStore store.Store // Initialize redisStore to avoid undefined error

	// Decide store based on config
	switch cfg.StoreType {
	case "postgres":
		postgresStore, err := postgres.New(cfg.PostgresDSN)
		if err != nil {
			logrus.Fatalf("Failed to connect to PostgreSQL: %v", err)
		}
		storeInstance = postgresStore
	case "redis":
		var err error
		redisStore, err = redis.New()
		if err != nil {
			logrus.Fatalf("Failed to connect to Redis: %v", err)
		}
		storeInstance = redisStore
	default:
		logrus.Fatal("Unsupported store type in config")
	}

	// Use in-memory store as fallback
	memoryStore := memory.New()

	hybridStore := &store.HybridStore{
		Postgres: storeInstance, // dynamic store selection
		Redis:    redisStore,
		Memory:   memoryStore,
	}

	urlService := service.NewURLService(hybridStore)

	// Middleware for request latency
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()
		requestLatency.WithLabelValues(c.Request.URL.Path).Observe(duration)
		logrus.WithFields(logrus.Fields{
			"path":    c.Request.URL.Path,
			"method":  c.Request.Method,
			"latency": duration,
			"status":  c.Writer.Status(),
		}).Info("Request processed")
	})

	router.SetupRouter(r, urlService)

	r.GET("/health", func(c *gin.Context) {
		if err := hybridStore.Health(); err != nil {
			c.JSON(500, gin.H{"status": "unhealthy", "error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "healthy"})
	})

	r.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})

	// Serve Swagger UI
	config := &ginSwagger.Config{
		URL: "http://localhost:8080/swagger/doc.json",
	}
	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	r.StaticFile("/swagger/doc.json", "./docs/swagger.json")

	if err := r.Run(":" + cfg.ServerPort); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}

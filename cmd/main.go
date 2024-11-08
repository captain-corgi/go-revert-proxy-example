package main

import (
	"net/http"
	"time"

	"github.com/captain-corgi/go-revert-proxy-example/config"
	"github.com/captain-corgi/go-revert-proxy-example/internal/cache"
	"github.com/captain-corgi/go-revert-proxy-example/internal/handlers"
	"github.com/captain-corgi/go-revert-proxy-example/internal/middleware"
	"github.com/captain-corgi/go-revert-proxy-example/internal/proxy"
	"github.com/captain-corgi/go-revert-proxy-example/internal/router"
	"github.com/captain-corgi/go-revert-proxy-example/pkg/logger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	log := logger.NewLogger()

	// Initialize Redis cache
	redisCache := cache.NewRedisCache(cfg.RedisAddr, cfg.CacheTTL)

	// Initialize reverse proxy
	reverseProxy := proxy.NewReverseProxy(cfg.Backends, log)

	// Initialize handler and router
	handler := handlers.NewProxyHandler(redisCache, reverseProxy, log)

	// Set up JWT middleware and rate limiter middleware
	jwtMiddleware := middleware.NewJWTMiddleware(cfg.JWTSecret, cfg.SkipAuthPaths)
	rateLimiter := middleware.NewRateLimiterMiddleware(5, time.Minute) // Example rate limit: 5 requests/minute

	// Initialize router with JWT and rate limiter middleware
	mux := router.NewRouter(handler, jwtMiddleware, rateLimiter)

	// Start the server
	log.Info("Starting server on " + cfg.ServerPort)
	srv := &http.Server{
		Addr:         cfg.ServerPort,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Error(err, "Server failed to start")
	}
}

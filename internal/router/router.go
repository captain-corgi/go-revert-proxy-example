package router

import (
	"net/http"

	"github.com/captain-corgi/go-revert-proxy-example/internal/handlers"
	"github.com/captain-corgi/go-revert-proxy-example/internal/middleware"
)

func NewRouter(handler *handlers.ProxyHandler, jwtMiddleware *middleware.JWTMiddleware, rateLimiter *middleware.RateLimiterMiddleware) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", rateLimiter.Middleware(jwtMiddleware.Middleware(handler)))
	return mux
}

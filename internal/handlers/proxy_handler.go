package handlers

import (
	"context"
	"net/http"

	"github.com/captain-corgi/go-revert-proxy-example/internal/cache"
	"github.com/captain-corgi/go-revert-proxy-example/internal/proxy"
	"github.com/captain-corgi/go-revert-proxy-example/pkg/logger"
)

type ProxyHandler struct {
	cache  *cache.RedisCache
	proxy  *proxy.ReverseProxy
	logger *logger.Logger
}

func NewProxyHandler(cache *cache.RedisCache, proxy *proxy.ReverseProxy, logger *logger.Logger) *ProxyHandler {
	return &ProxyHandler{cache: cache, proxy: proxy, logger: logger}
}

func (h *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cacheKey := r.URL.String()

	if r.Method == http.MethodGet {
		if cachedResponse, found := h.cache.Get(ctx, cacheKey); found {
			h.logger.Info("Serving cached response for " + cacheKey)
			w.Write(cachedResponse)
			return
		}
	}

	rec := proxy.NewResponseRecorder(w)
	h.proxy.ServeHTTP(rec, r)

	if r.Method == http.MethodGet && rec.StatusCode == http.StatusOK {
		if err := h.cache.Set(ctx, cacheKey, rec.Body.Bytes()); err != nil {
			h.logger.Error(err, "Failed to cache response for "+cacheKey)
		} else {
			h.logger.Info("Cached response for " + cacheKey)
		}
	}
}

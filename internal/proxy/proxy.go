package proxy

import (
	"bytes"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/captain-corgi/go-revert-proxy-example/pkg/logger"
)

type ReverseProxy struct {
	backends []string
	mu       sync.Mutex
	counter  int
	logger   *logger.Logger
}

func NewReverseProxy(backends []string, logger *logger.Logger) *ReverseProxy {
	return &ReverseProxy{backends: backends, logger: logger}
}

func (p *ReverseProxy) nextBackend() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	server := p.backends[p.counter]
	p.counter = (p.counter + 1) % len(p.backends)
	p.logger.Info("Forwarding request to backend " + server)
	return server
}

func (p *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	targetURL, _ := url.Parse(p.nextBackend())
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		p.logger.Error(err, "Error forwarding request")
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
	}

	proxy.ServeHTTP(w, r)
}

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       *bytes.Buffer
}

func NewResponseRecorder(w http.ResponseWriter) *ResponseRecorder {
	return &ResponseRecorder{ResponseWriter: w, Body: new(bytes.Buffer)}
}

func (r *ResponseRecorder) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *ResponseRecorder) Write(b []byte) (int, error) {
	r.Body.Write(b)
	return r.ResponseWriter.Write(b)
}

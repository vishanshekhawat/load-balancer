package srv

import (
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type ServerI interface {
	GetActiveConnection() int64
	IsHealthy() bool
}

type Server struct {
	URL              *url.URL
	ActiveConnection atomic.Int64
	HealthStatus     bool
}

func New(url *url.URL) *Server {

	return &Server{
		URL: url,
	}
}

// GetActiveConnection
func (srv *Server) GetActiveConnection() int64 {
	return srv.ActiveConnection.Load()
}

// AddConnectionCount
func (srv *Server) AddConnectionCount() {
	srv.ActiveConnection.Add(1)
}

// SubstractConnectionCount
func (srv *Server) SubstractConnectionCount() {
	srv.ActiveConnection.Add(-1)
}

// IsHealthy
func (srv *Server) IsHealthy() bool {
	return srv.HealthStatus
}

func (srv *Server) Proxy() *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(srv.URL)
}

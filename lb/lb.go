package lb

import (
	"lb/srv"
	"sync"
)

type LoadBalancers struct {
	Srvs       []srv.Server
	RoundRobin int
	mu         sync.Mutex
}

func NewLoadBalancer(servers []srv.Server) *LoadBalancers {
	return &LoadBalancers{
		Srvs:       servers,
		RoundRobin: -1,
	}
}

func (lb *LoadBalancers) NextServerLeastConnection() *srv.Server {
	nextSrv := &lb.Srvs[0]
	leastConnections := nextSrv.GetActiveConnection()
	for i := 1; i < len(lb.Srvs); i++ {
		connCount := lb.Srvs[i].GetActiveConnection()
		if connCount < int64(leastConnections) {
			nextSrv = &lb.Srvs[i]
		}
	}
	return nextSrv
}

func (lb *LoadBalancers) NextServerRoundRobin() *srv.Server {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	totalServers := len(lb.Srvs)
	if lb.RoundRobin == totalServers {
		lb.RoundRobin = 0
		return &lb.Srvs[0]
	}

	nextSrv := &lb.Srvs[lb.RoundRobin]
	lb.RoundRobin += 1

	return nextSrv
}

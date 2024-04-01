package lb

import "lb/srv"

type LoadBalancers struct {
	Srvs []srv.Server
}

func NewLoadBalancer(servers []srv.Server) *LoadBalancers {
	return &LoadBalancers{
		Srvs: servers,
	}
}

func (lb *LoadBalancers) NextServerLeastActive() *srv.Server {
	nextSrv := &lb.Srvs[0]
	leastConnections := 0
	for i := 1; i < len(lb.Srvs); i++ {
		connCount := lb.Srvs[i].GetActiveConnection()
		if connCount < int64(leastConnections) {
			nextSrv = &lb.Srvs[i]
		}
	}
	return nextSrv
}

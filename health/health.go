package health

import "net"

type HealthCheck struct {
	Server string
	Status string
}

func (h *HealthCheck) CheckHealth() {

	_, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		h.Status = "DOWN"
		return
	}

}

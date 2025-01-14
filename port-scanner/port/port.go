package port

import (
	"net"
	"strconv"
	"time"
)

type PortStatus struct {
	Port int
	Open bool
}

func ScanPorts(host string) []PortStatus {
	var result []PortStatus

	for i := 1; i <= 65535; i++ {
		address := host + ":" + strconv.Itoa(i)
		_, err := net.DialTimeout("tcp", address, 5*time.Second)
		result = append(result, PortStatus{Port: i, Open: err == nil})
	}

	return result
}

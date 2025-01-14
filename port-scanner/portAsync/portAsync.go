package portAsync

import (
	"net"
	"strconv"
	"sync"
	"time"
)

type PortStatus struct {
	Port int
	Open bool
}

func ScanPorts(host string) []PortStatus {
	const tcpPorts = 65535
	ch := make(chan struct{}, 100)

	result := make([]PortStatus, tcpPorts)

	var wg sync.WaitGroup

	for i := 1; i <= tcpPorts; i++ {
		ch <- struct{}{}

		wg.Add(1)
		go scanPort(host, i, &result, &wg, ch)
	}

	wg.Wait()

	return result
}

func scanPort(host string, port int, result *[]PortStatus, wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	defer func() { <-ch }()

	address := host + ":" + strconv.Itoa(port)
	_, err := net.DialTimeout("tcp", address, 5*time.Second)
	(*result)[port-1] = PortStatus{Port: port, Open: err == nil}
}

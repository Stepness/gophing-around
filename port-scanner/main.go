package main

import (
	"fmt"
	"github.com/stepness/gophing-around/port-scanner/port"
	"github.com/stepness/gophing-around/port-scanner/portAsync"
	"time"
)

func main() {
	start := time.Now()
	_ = port.ScanPorts("127.0.0.1")
	duration := time.Since(start)

	fmt.Printf("Sync - It took %d ms\n", duration.Milliseconds())

	startConcurrent := time.Now()
	portsAsync := portAsync.ScanPorts("127.0.0.1")
	durationConcurrent := time.Since(startConcurrent)

	fmt.Printf("Async - It took %d ms\n", durationConcurrent.Milliseconds())

	if len(portsAsync) == 0 {
		fmt.Println("No ports open")
	}

	fmt.Print("\nOpen ports\n")
	for _, p := range portsAsync {
		if p.Open {
			fmt.Printf("Port %d - Open\n", p.Port)
		}
	}
}

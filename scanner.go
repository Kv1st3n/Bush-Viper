package main

import (
	"net"
	"strconv"
	"sync"
	"time"
)

type ScanResult struct {
	IP     string
	Port   string
	Opened string
}

func singlePortScan(address string, port string) (ScanResult, error) {

	singlePort := net.JoinHostPort(address, port)

	conn, err := net.DialTimeout("tcp", singlePort, 300*time.Millisecond)

	if err != nil {
		return ScanResult{}, nil
	}
	defer conn.Close()

	return ScanResult{
		IP:     address,
		Port:   port,
		Opened: "Open",
	}, nil

}

func widePortScan(address string) ([]ScanResult, error) {

	var mu sync.Mutex
	var wg sync.WaitGroup
	var semGroup = make(chan struct{}, 100)

	var results []ScanResult

	for startPort := 1; startPort <= 65536; startPort++ {
		wg.Add(1)
		semGroup <- struct{}{}

		go func(id int) {
			defer wg.Done()
			defer func() {
				<-semGroup
			}()

			port := strconv.Itoa(startPort)

			target := net.JoinHostPort(address, port)

			conn, err := net.DialTimeout("tcp", target, 400*time.Millisecond)
			if err != nil {
				return
			}
			defer conn.Close()

			mu.Lock()
			results = append(results, ScanResult{
				IP:     address,
				Port:   port,
				Opened: "Opened",
			})
			mu.Unlock()

		}(startPort)

	}

	wg.Wait()
	return results, nil

}

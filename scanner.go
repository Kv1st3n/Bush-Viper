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
	Banner string
}

func singlePortScan(address string, port string) (ScanResult, error) {

	singlePort := net.JoinHostPort(address, port)

	conn, err := net.DialTimeout("tcp", singlePort, 500*time.Millisecond)

	if err != nil {
		return ScanResult{}, nil
	}
	defer conn.Close()

	portInt, _ := strconv.Atoi(port)

	banner, _ := grabBanner(conn, portInt)

	return ScanResult{
		IP:     address,
		Port:   port,
		Opened: "Open",
		Banner: banner,
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

			banner, _ := grabBanner(conn, id)

			conn.Close()

			mu.Lock()
			results = append(results, ScanResult{
				IP:     address,
				Port:   port,
				Opened: "Opened",
				Banner: banner,
			})
			mu.Unlock()

		}(startPort)

	}

	wg.Wait()
	return results, nil

}

func grabBanner(conn net.Conn, port int) (string, error) {

	buffer := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	n, err := conn.Read(buffer)
	if err == nil && n > 0 {
		return string(buffer[:n]), nil
	}

	service, exists := ServiceDB[port]
	if exists && len(service.Probe) > 0 {
		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		_, _ = conn.Write(service.Probe)

		conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		n, err = conn.Read(buffer)
		if err == nil && n > 0 {
			return string(buffer[:n]), nil
		}
	}
	return "", nil

}

func isPortInBanner(currentPort int) bool {

	service, exists := ServiceDB[currentPort]

	if !exists {
		return false
	}

	return len(service.Probe) > 0
}

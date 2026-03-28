package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

// to hold data for successful port connection
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

	// port is open, check if there is a banner probe for it
	banner, _ := grabBanner(conn, portInt)

	return ScanResult{
		IP:     address,
		Port:   port,
		Opened: "Open",
		Banner: banner,
	}, nil

}

func portWorkerPool(adress string, ports <-chan int, results chan<- ScanResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for p := range ports {
		portStr := strconv.Itoa(p)
		targetAdress := net.JoinHostPort(adress, portStr)

		conn, err := net.DialTimeout("tcp", targetAdress, 400*time.Millisecond)
		if err != nil {
			continue
		}
		// port is open, check if there is a banner probe for it
		banner, _ := grabBanner(conn, p)
		conn.Close()

		results <- ScanResult{
			IP:     adress,
			Port:   portStr,
			Opened: "Opened",
			Banner: banner,
		}

	}
}

func widePortScan(address string, start string, end string) ([]ScanResult, error) {

	// convert startP to int
	startP, err := strconv.Atoi(start)
	if err != nil {
		return nil, fmt.Errorf("invalid start port '%s': %w", start, err)
	}

	// convert startP to int
	endP, err := strconv.Atoi(end)
	if err != nil {
		return nil, fmt.Errorf("invalid end port '%s': %w", end, err)
	}

	// ensure that the startp is smaller than the endport for the range
	if startP > endP {
		return nil, fmt.Errorf("start port (%d) cannot be greater than end port (%d)", startP, endP)
	}

	portQueue := make(chan int, 100)
	resultChan := make(chan ScanResult, 100)

	var wg sync.WaitGroup
	const portWorkerCount = 100

	for i := 0; i < portWorkerCount; i++ {
		wg.Add(1)
		go portWorkerPool(address, portQueue, resultChan, &wg)
	}

	go func() {
		for p := startP; p <= endP; p++ {
			portQueue <- p
		}
		close(portQueue)
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var results []ScanResult
	for res := range resultChan {
		results = append(results, res)
	}

	return results, nil

}

func grabBanner(conn net.Conn, port int) (string, error) {

	buffer := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	// read initially in case the service talks first
	n, err := conn.Read(buffer)
	if err == nil && n > 0 {
		return string(buffer[:n]), nil
	}

	// if service does not talk
	// check if a service exists based on the targeted port, e.g. HTTP for their probe
	if !isPortInBanner(port) {
		return "", nil
	}

	// formulate a new probe based on the service
	service := ServiceDB[port]
	_, err = conn.Write(service.Probe)
	if err != nil {
		return "", err
	}

	conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	// read the response
	n, err = conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil

}

// checks if the targeted port is in the db, and if it has some form of probe
func isPortInBanner(currentPort int) bool {

	service, exists := ServiceDB[currentPort]

	if !exists {
		return false
	}

	return len(service.Probe) > 0
}

package main

import (
	"net"
	"strconv"
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

	var results []ScanResult

	// 5000 for now, will use goroutines later
	for startPort := 1; startPort <= 5000; startPort++ {

		port := strconv.Itoa(startPort)

		target := net.JoinHostPort(address, port)

		conn, err := net.DialTimeout("tcp", target, 400*time.Millisecond)
		if err != nil {
			continue
		}
		defer conn.Close()

		results = append(results, ScanResult{
			IP:     address,
			Port:   port,
			Opened: "Opened",
		})

	}

	return results, nil

}

package main

import (
	"net"
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

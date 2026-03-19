package main

import (
	"net"
)

func getHostAddress(hostName string) ([]string, error) {

	ips, err := net.LookupHost(hostName)
	if err != nil {
		return nil, err
	}

	return ips, err
}

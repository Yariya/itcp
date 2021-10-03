package main

import "net"

func GetHostname(addr string) (string, error) {
	hostname, err := net.LookupAddr(addr)
	if err != nil {
		return "", err
	}
	var Hostname string
	for _, name := range hostname {
		Hostname = name
	}

	return Hostname, nil
}


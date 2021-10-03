package main

import (
	"net"
	"time"
)

func PingTCP(address string, port string) (time.Duration, error){
	for {
		t := time.Now()
		s, err := net.DialTimeout("tcp", address+":"+port, time.Millisecond*time.Duration(socktimeout))
		if err != nil {
			return 0, err
		}
		s.Close()
		return time.Since(t), err
	}
}

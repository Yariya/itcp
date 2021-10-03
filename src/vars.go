package main

var PingStop = make(chan struct{})


var CommandFlags = []string {
	"--notcp", // Only ICMP Ping
	"--noicmp", // Only TCP Ping
	"--de", // Deactivates ISP, Hostname Feature
}

var (
	IcmpSuccess int8
	TcpSuccess int8
)
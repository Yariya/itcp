package main

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"os"
	"time"
)

func PingICMP(address string) (time.Duration, error) {

	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0") // 0.0.0.0 Default
	if err != nil {
		return 0, err
	}
	defer c.Close()

	dst, err := net.ResolveIPAddr("ip4", address)
	if err != nil {
		return 0, err
	}

	m := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("iTCP Pinger (c) Yariya"),
		},
	}
	b, err := m.Marshal(nil)
	if err != nil {
		return 0, err
	}

	t := time.Now()
	x, err := c.WriteTo(b, dst)
	if err != nil {
		return 0, err
	} else if x != len(b) {
		return 0, fmt.Errorf("ERR [%v] searched %v", x, len(b))
	}

	reply := make([]byte, 1500)

	err = c.SetReadDeadline(time.Now().Add(time.Second*1))
	if err != nil {
		return 0, err
	}

	n, _, err := c.ReadFrom(reply)
	if err != nil {
		return 0, err
	}

	end := time.Since(t)

	parsed, err := icmp.ParseMessage(1, reply[:n]) // Protocol 1 == ICMP!!!!
	if err != nil {
		return 0, err
	}

	switch parsed.Type {
	case ipv4.ICMPTypeEchoReply:
		return end, nil
	default:
		return 0, fmt.Errorf("Invalid Echo reply!")
	}

}

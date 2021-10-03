package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)


var (
	DeactivateFlag = false
	NoTCP = false
	NoICMP = false
)

var timeout = 1000
var socktimeout = 1000
var threads = 1

var Isp string

func iTCP() {
	Cls()
	addr := os.Args[1]
	port := os.Args[3]

	/*
	Hostname, err := GetHostname(addr)
	if err != nil {
		fmt.Println("Cant resolve hostname!")
		return
	}
	 */
	CheckFlags()

	if Isp == "" {
		is, err := ISP(addr)
		if err != nil {
			panic(err)
		}
		Isp = is
	}


	for {
		select {
		case <-PingStop:
			return
		default:


			if NoTCP {
				var IcmpErr bool
				icmp, err := PingICMP(addr)
				if err != nil {
					IcmpErr = true
				}
				if DeactivateFlag {
					if !IcmpErr{
						fmt.Printf("\033[37mConnected to \u001B[32m%s\u001B[37m: icmp=\u001B[32m%.5sms\u001B[37m port=\u001B[32m%s\u001B[37m\n", addr, icmp, port)
					} else {
						fmt.Println("\033[31mConnection timed out\033[39m")
					}
				} else {
					if !IcmpErr{
						fmt.Printf("\033[37mConnected to \u001B[32m%s (%s)\u001B[37m: icmp=\u001B[32m%.5sms\u001B[37m port=\u001B[32m%s\u001B[37m\n", addr, Isp, icmp, port)
					} else {
						fmt.Println("\033[31mConnection timed out\033[39m")
					}
				}
			} else if NoICMP {
				var TcpErr bool
				tcp, err := PingTCP(addr, port)
				if err != nil {
					TcpErr = true
				}

				if DeactivateFlag {
					if !TcpErr{
						fmt.Printf("\033[37mConnected to \u001B[32m%s\u001B[37m: tcp=\u001B[32m%.5sms\u001B[37m port=\u001B[32m%s\u001B[37m\n", addr, tcp, port)
					} else {
						fmt.Println("\033[31mConnection timed out\033[39m")
					}
				} else {
					if !TcpErr{
						fmt.Printf("\033[37mConnected to \u001B[32m%s (%s)\u001B[37m: tcp=\u001B[32m%.5sms\u001B[37m port=\u001B[32m%s\u001B[37m\n", addr, Isp, tcp, port)
					} else {
						fmt.Println("\033[31mConnection timed out\033[39m")
					}
				}


			} else {
				var IcmpErr bool
				var TcpErr bool
				icmp, err := PingICMP(addr)
				if err != nil {
					IcmpErr = true
				}
				tcp, err := PingTCP(addr, port)
				if err != nil {
					TcpErr = true
				}

				if DeactivateFlag {
					if !IcmpErr && !TcpErr {
						fmt.Printf("\033[37mConnected to \u001B[32m%s\u001B[37m: tcp=\u001B[32m%.5sms\u001B[37m icmp=\u001B[32m%.5sms\u001B[37m port=\u001B[32m%s\u001B[37m\n", addr, tcp, icmp, port)
					} else {
						if TcpErr && IcmpErr{
							fmt.Println("\033[31mConnection timed out\033[39m")
						} else if IcmpErr == true && TcpErr != false {
							fmt.Printf("\033[37mConnected to \u001B[32m%s\u001B[37m: tcp=\u001B[32m%.5s\u001B[37m icmp=\u001B[31mTIMEOUT\u001B[37m port=\u001B[32m%s\u001B[37m\n", addr,  tcp, port)
						} else if TcpErr == true && IcmpErr == false{
							fmt.Printf("\033[37mConnected to \u001B[32m%s\u001B[37m: tcp=\u001B[31mTIMEOUT\u001B[37m icmp=\u001B[32m%.5sms\u001B[37m port=\u001B[32m%s\u001B[37m\n", addr, icmp, port)
						} else if TcpErr && IcmpErr{
							fmt.Println("\033[31mConnection timed out\033[39m")
						}
					}
				} else {
					if !IcmpErr && !TcpErr {
						fmt.Printf("\033[37mConnected to \u001B[32m%s (%s)\u001B[37m: tcp=\u001B[32m%.5sms\u001B[37m icmp=\u001B[32m%.5sms\u001B[37m port=\u001B[32m%s\u001B[37m\n", addr, Isp, tcp, icmp, port)
					} else {
						if TcpErr && IcmpErr{
							fmt.Println("\033[31mConnection timed out\033[39m")
						} else if IcmpErr == true && TcpErr != false {
							fmt.Printf("\033[37mConnected to \u001B[32m%s (%s)\u001B[37m: tcp=\u001B[32m%.5s\u001B[37m icmp=\u001B[31mTIMEOUT\u001B[37m port=\u001B[32m%s\u001B[37m\n", addr,  Isp, tcp, port)
						} else if TcpErr == true && IcmpErr == false{
							fmt.Printf("\033[37mConnected to \u001B[32m%s (%s)\u001B[37m: tcp=\u001B[31mTIMEOUT\u001B[37m icmp=\u001B[32m%.5sms\u001B[37m port=\u001B[32m%s\u001B[37m\n", addr,  Isp, icmp, port)
						} else if TcpErr && IcmpErr{
							fmt.Println("\033[31mConnection timed out\033[39m")
						}
					}
				}
			}

			time.Sleep(time.Millisecond*time.Duration(timeout))
		}
	}
}
func main() {
	if len(os.Args) < 4 {
		help(false)
		return
	}
	if os.Args[2] != "-p"{help(false);return}

	t, s, thrs, err := ValidateTimeout()
	if err != nil {
		help(false)
		return
	}

	if t != 0 {
		timeout=t
	}
	if s != 0 {
		socktimeout=s
	}
	if thrs != 0 {
		threads = thrs
	}

	/*if t != 0 {
		timeout=t
	} else if s != 0 {
		socktimeout=s
	} else if thrs != 0 {
		threads = thrs
	} else if s != 0 && t != 0 {
		timeout=t
		socktimeout=s
	} else if s != 0 && t != 0 && thrs != 0{
		timeout=t
		socktimeout=s
		threads=thrs
	}
	 */


	for x := 0; x < threads; x++ {
		go iTCP()
		time.Sleep(time.Millisecond*50) // Default
	}
	callback := make(chan os.Signal, 1)
	signal.Notify(callback, os.Kill, os.Interrupt)
	<-callback
	close(PingStop)
	<-PingStop

	fmt.Println(" 			iTCP Stopped STRG^C 		")
}

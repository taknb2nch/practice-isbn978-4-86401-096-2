package main

import (
	"fmt"
	"net"
	"strconv"
)

func tryConnect(network, host string, port int) net.Conn {
	p := strconv.Itoa(port)
	addr := net.JoinHostPort(host, p)
	c, e := net.Dial(network, addr)

	if e == nil {
		return c
	}

	return nil
}

func connect(network, service, host string) net.Conn {
	_, addrs, _ := net.LookupSRV(service, network, host)

	for _, srv := range addrs {
		fmt.Printf("net.LookupSRV: %s, %s, %d\n", network, srv.Target, srv.Port)

		c := tryConnect(network, srv.Target, int(srv.Port))

		if c != nil {
			return c
		}
	}

	port, _ := net.LookupPort(network, service)
	ips, _ := net.LookupHost(host)

	for _, ip := range ips {
		fmt.Printf("net.LookupHost: %s, %s, %d\n", network, ip, port)

		c := tryConnect(network, ip, port)

		if c != nil {
			return c
		}
	}

	return nil
}

func main() {
	c := connect("tcp", "http", "informit.com")
	c.Write([]byte("GET / HTTP/1.1\r\nHost: informit.com\r\n\r\n"))

	buffer := make([]byte, 1024)

	c.Read(buffer)

	fmt.Printf("%s", buffer)
}

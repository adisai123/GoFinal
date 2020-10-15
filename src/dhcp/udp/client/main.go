package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"syscall"
)

func main() {
	p := make([]byte, 2048)
	dialer := &net.Dialer{
		Control: func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				err := syscall.SetsockoptInt(int(fd), syscall.IPPROTO_IP, syscall.IP_TOS, 123)
				if err != nil {
					log.Printf("control: %s", err)
					return
				}
			})
		},
	}
	conn, err := dialer.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hi UDP Server, How are you doings?")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}

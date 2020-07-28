package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"syscall"
	"time"
)

//https://stackoverflow.com/questions/40544096/how-to-set-socket-option-ip-tos-for-http-client-in-go-language
func main() {
	dialer := &net.Dialer{
		Control: func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				err := syscall.SetsockoptInt(int(fd), syscall.IPPROTO_IP, syscall.IP_TOS, 128)
				if err != nil {
					log.Printf("control: %s", err)
					return
				}
			})
		},
	}
	tr := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	s := &http.Client{
		Transport: tr,
	}
	fmt.Println(s)
	resp, err := s.Get("http://192.168.43.5:8081/users")
	if err != nil {
		log.Fatalf("GET error: %v", err)
	}

	log.Printf("got %q", resp.Status)
}

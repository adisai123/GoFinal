package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"syscall"
	"time"
)

func main() {
	dial := func(ctx context.Context, network, addr string) (net.Conn, error) {
		conn, err := (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext(ctx, network, addr)
		if err != nil {
			return nil, err
		}

		tcpConn, ok := conn.(*net.TCPConn)
		if !ok {
			err = errors.New("conn is not tcp")
			return nil, err
		}

		f, err := tcpConn.File()
		if err != nil {
			return nil, err
		}

		err = syscall.SetsockoptInt(int(f.Fd()), syscall.IPPROTO_IP, syscall.IP_TOS, int(435))
		if err != nil {
			return nil, err
		}

		return conn, nil
	}
	tr := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dial,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	s := &http.Client{
		Transport: tr,
	}
	fmt.Println(s)
	resp, err := s.Get("http://169.254.220.69:8080/")
	if err != nil {
		log.Fatalf("GET error: %v", err)
	}

	log.Printf("got %q", resp.Status)
}

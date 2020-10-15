package main

import (
	"context"
	"errors"
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

		err = syscall.SetsockoptInt(int(f.Fd()), syscall.IPPROTO_IP, syscall.IP_TOS, 128)
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
	c := &http.Client{
		Transport: tr,
	}
	resp, err := c.Get("http://192.168.43.5:8080/user")
	if err != nil {
		log.Fatalf("GET error: %v", err)
	}
	body := []byte{}
	resp.Body.Read(body)
	log.Printf("got %s", string(body))
}

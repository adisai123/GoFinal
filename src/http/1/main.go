package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"syscall"
	"time"
)

func main() {
	var i i = 10
	http.ListenAndServe(":8080", i)
}

type i int

//handler interface for http
func (i i) ServeHTTP(rs http.ResponseWriter, r *http.Request) {
	//rs.Write([]byte("hi"))
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

	fmt.Fprintln(rs, "another way to responde")
}

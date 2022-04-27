package NetUtil

import (
	"fmt"
	"net"
)

func Listener(ip string, port int64, msg chan string) {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
		}
		go func(net.Conn) {
			buf := make([]byte, 1024)
			_, err = conn.Read(buf)
			if err != nil {
			}
			for len(msg) > 0 {
				<-msg
			}
			msg <- fmt.Sprintf("%s", buf)
			conn.Write([]byte("Message received."))
			conn.Close()
		}(conn)

	}
}

func Send(ip string, port int64, msg string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return
	}
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return
	}
	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		return
	}

	conn.Close()
}

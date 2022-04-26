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
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
		}
		msg <- string(buf)
		conn.Write([]byte("Message received."))
		conn.Close()

	}
}

func Send(ip, msg string, port int64) {
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

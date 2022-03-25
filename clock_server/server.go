package clock_server

import (
	"io"
	"log"
	"net"
	"time"
)

func Run(protocol string, port string, isSequential bool) {
	listener, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		if isSequential {
			handleConn(conn)
		} else {
			go handleConn(conn)
		}
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

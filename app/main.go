package main

import (
	"fmt"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func handleClient(conn net.Conn) {
	
	defer conn.Close()

	buff := make([]byte, 1024)

	for {
		_, err := conn.Read(buff)
		if err != nil {
			fmt.Println("error reading data", err.Error())
			break
		}
		msg, err := conn.Write([]byte("+PONG\r\n"))
	}
	conn.Close()
}

func main() {

	// listens on port 6379 and exits with an error if the server fails to start
	listener, err := net.Listen("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("error starting server: ", err.Error())
		os.Exit(1)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err.Error())
			continue
		}
		
		handleClient(conn)
	}

}

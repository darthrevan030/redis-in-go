package main

import (
	"fmt"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

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
			fmt.Println("error accepting connection:", err.Error())
			continue
		}
	go handleClient(conn)
	}

}

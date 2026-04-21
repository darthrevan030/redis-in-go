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

	resp := NewResp(conn)

	for {
		value, err := resp.Read()
		if err != nil {
			fmt.Println("error in resp", err.Error())
			break
		}

		fmt.Println(value)

		if value.typ == "array" && len(value.array) > 0 {
			cmd := value.array[0].bulk
			if cmd == "ECHO" && len(value.array) >= 1 {
				arg := value.array[1].bulk
				response := fmt.Sprintf("%d\r\n%s\r\n", len(arg), arg)
				conn.Write([]byte(response))
			} else {
				conn.Write([]byte("+PONG\r\n"))
			}
		} else {
			conn.Write([]byte("+PONG\r\n"))
		}
	}
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
			fmt.Println("error accepting connection:", err.Error())
			continue
		}
	go handleClient(conn)
	}



}

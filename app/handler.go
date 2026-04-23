package main

import (
	"net"
	"fmt"

)

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
			if cmd == "ECHO" && len(value.array) >= 2 {
				arg := value.array[1].bulk
				response := fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
				conn.Write([]byte(response))
			} else {
				cmd := value.array[0].bulk
				if cmd == "SET" && len(value.array) >= 2 {
					arg := value.array[1].bulk
					response := fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
					conn.Write([]byte(response))
				}
			}
		} else {
			conn.Write([]byte("+PONG\r\n"))
		}
	}
}



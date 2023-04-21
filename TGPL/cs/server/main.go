package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("The server is running...")
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Failed to listen")
		return
	}
	conn, err1 := listener.Accept()

	for {
		if err1 != nil {
			fmt.Println("Failed to connect")
		} else {
			fmt.Println("Connected successfully... conn", conn)
			fmt.Println("Message:", conn.RemoteAddr().String())
		}
	}
}

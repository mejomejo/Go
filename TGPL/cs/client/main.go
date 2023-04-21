package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("The client is running...")

	coon, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Failed to connect...")
		return
	}
	fmt.Println("Connected successfully... conn:", coon)
}

package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器8889端口监听......")
	l, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("Listen()", err)
		return
	}

	defer l.Close()

	for {
		c, err2 := l.Accept()
		if err2 != nil {
			fmt.Println("Accept()", err2)
		}

		go process(c)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024*4)
		n, err := conn.Read(buf[:])
		if n != 4 || err != nil {
			fmt.Println("Read()", err)
			return
		}

		fmt.Println(buf[:4])

	}
}

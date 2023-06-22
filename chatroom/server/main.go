package main

import (
	"chatroom/common/utils"
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:2004")
	if err != nil {
		return
	}
	defer l.Close()

	for {
		c, err2 := l.Accept()
		if err2 != nil {
			return
		}

		go handle(c)
	}
}

func handle(c net.Conn) {

	c.Write([]byte("You are connected..."))

	for {
		err := utils.Recv(c)
		if err != nil {
			if err == io.EOF {
				fmt.Println("disconnected")
				break
			}
			fmt.Println(err)
			break
		}
	}
}

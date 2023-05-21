package main

import (
	"chatroom/c/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("Listen()", err)
		return
	}

	defer l.Close()
	fmt.Println("服务器8889端口监听......")

	for {
		c, err2 := l.Accept()
		if err2 != nil {
			fmt.Println("Accept()", err2)
		}

		go process(c)
	}
}

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 1024*4)
	n, err := conn.Read(buf[:])
	if n != 4 || err != nil {
		//err = errors.New("read pkg head err")
		return
	}

	fmt.Println(buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	n2, err := conn.Read(buf[:pkgLen])
	if n2 != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body err")
		return
	}

	//反序列化

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		return
	}

	return
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端正常退出")
				return
			}
			fmt.Println(err)
			return
		}

		fmt.Println("mes:", mes)

	}
}

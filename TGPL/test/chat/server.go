package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	ip   string
	port int

	Onlinemap map[string]*User
	mapLock   sync.Mutex

	Message chan string
}

func NewServer(ip string, port int) (t *Server) {
	t = &Server{
		ip:        ip,
		port:      port,
		Onlinemap: make(map[string]*User),
		Message:   make(chan string),
	}
	return
}

func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		for _, user := range this.Onlinemap {
			user.C <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) BroadCast(user *User, msg string) {
	send := "[" + user.Addr + "]" + user.Name + msg
	this.Message <- send
}

func (this *Server) handler(conn net.Conn) {
	user := NewUser(conn, this)
	user.Online()
	go func() {
		buf := make([]byte, 1024*3)
		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println(err)
				return
			}
			if n == 0 {
				user.Offline()
				return
			}
			msg := string(buf[:n-1])

			user.DoMessage(msg)
		}
	}()
}

func (this *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.ip, this.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	go this.ListenMessager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go this.handler(conn)

	}
}

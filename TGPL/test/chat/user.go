package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	addr := conn.RemoteAddr().String()
	t := &User{
		Name:   addr,
		Addr:   addr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	go t.ListenMessage()

	return t
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

func (this *User) Online() {
	this.server.mapLock.Lock()
	this.server.Onlinemap[this.Name] = this
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "已上线")
}

func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.Onlinemap, this.Name)
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "已下线")
}

func (this *User) DoMessage(msg string) {
	if msg == "/all people" {
		this.server.mapLock.Lock()
		for _, user := range this.server.Onlinemap {
			send := ("[" + user.Addr + "]" + user.Name + ":在线\n")
			this.SendMsg(send)
		}
		this.server.mapLock.Unlock()
		return
	}
	if strings.HasPrefix(msg, "/rename ") {
		newname := strings.Split(msg, " ")[1]
		_, ok := this.server.Onlinemap[newname]
		if ok {
			this.SendMsg("该用户名已经被使用\n")
			return
		}
		this.server.mapLock.Lock()
		delete(this.server.Onlinemap, this.Name)
		this.server.Onlinemap[newname] = this
		this.server.mapLock.Unlock()

		this.Name = newname
		return
	}
	if strings.HasPrefix(msg, "/@") {
		val := strings.Split(msg, "@")[1]
		name := strings.Split(val, " ")[0]
		remoteuser, ok := this.server.Onlinemap[name]
		if !ok {
			this.SendMsg("该用户不存在")
			return
		}
		remoteuser.SendMsg(this.Name + "对您说" + strings.Split(val, " ")[1])
		return
	}

	this.server.BroadCast(this, msg)
}

func (this *User) ListenMessage() {

	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

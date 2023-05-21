package main

import (
	"chatroom/c/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func login(userId int, userPwd string) (err error) {
	// fmt.Println(userId, userPwd)
	// fmt.Println("登录成功")

	// return nil

	c, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("Dial()", err)
		return
	}

	defer c.Close()

	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("Marshal", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("Marshal", err)
		return
	}

	//向服务器发送
	var pkgLen uint32 = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:], pkgLen)

	//发送长度
	n, err := c.Write(bytes[:])
	if n != 4 || err != nil {
		fmt.Println("Write()", err)
		return
	}

	fmt.Printf("客户端发送成功\n内容：%s\n长度：%d\n", string(data), len(data))

	_, err = c.Write(data)
	if err != nil {
		fmt.Println("Write()", err)
		return
	}

	time.Sleep(5 * time.Second)

	return
}

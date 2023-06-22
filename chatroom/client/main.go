package main

import (
	"chatroom/common"
	"chatroom/common/message"
	"chatroom/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	// 连接服务器
	conn, err := net.Dial("tcp", "localhost:2004")
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}
	defer conn.Close()

	go recvServer(conn)

	mes := message.Message{
		MessageType: message.UserMessage,
		Data:        newStudent(),
	}

	for i := 0; i < 10; i++ {
		utils.SendMsg(conn, mes)
	}

}

func newStudent() []byte {
	stu := common.Student{
		Name:  "党宇鑫",
		Sex:   "男",
		Id:    "04221145",
		Grade: "计科2204",
	}

	b, err2 := json.Marshal(stu)
	if err2 != nil {
		fmt.Println(err2)
		return nil
	}
	return b
}

func recvServer(c net.Conn) {
	buf := make([]byte, 4096)
	for {
		n, _ := c.Read(buf)
		if n <= 0 {
			continue
		}
		fmt.Println(string(buf[:n]))
	}
}

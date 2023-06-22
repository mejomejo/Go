package utils

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strconv"
)

func Send[T string | []byte](c net.Conn, data T) {
	send := fmt.Sprintf("%08d%s", len(data), data)
	c.Write([]byte(send))
}

func SendMsg(c net.Conn, mes message.Message) {
	switch mes.MessageType {
	case message.UserMessage:
		mesString, _ := json.Marshal(mes)
		Send(c, mesString)
	}
}

func Recv(c net.Conn) error {
	mes := message.Message{}
	read, err := RecvMsg(c)
	if err != nil {
		return err
	}
	json.Unmarshal(read, &mes)

	switch mes.MessageType {
	case message.UserMessage:
		fmt.Println(string(mes.Data))
	}

	return nil
}
func RecvMsg(c net.Conn) ([]byte, error) {
	l := make([]byte, 8)
	_, err := io.ReadFull(c, l)
	if err != nil {
		return nil, err
	}
	length, _ := strconv.Atoi(string(l))
	if length == 0 {
		return nil, io.EOF
	}
	dataBuf := make([]byte, length)
	_, err = io.ReadFull(c, dataBuf)
	if err != nil {
		return nil, err
	}

	return dataBuf, nil
}

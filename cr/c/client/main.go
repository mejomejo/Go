package main

import (
	"fmt"
)

var (
	userId  int
	userPwd string
)

func main() {
	loop := true
	var choice int

	for loop {
		fmt.Println("\t\t\t1.登录")
		fmt.Println("\t\t\t2.注册")
		fmt.Println("\t\t\t3.退出")

		fmt.Scanf("%d\n", &choice)
		loop = false

		switch choice {
		case 1:
			loop = false
		case 2:
			fmt.Println("开始注册")
		case 3:
			fmt.Println("退出成功")
		default:
			fmt.Println("您的输入有误，请重试")
			loop = true
		}
	}

	if choice == 1 {
		fmt.Println("请输入用户id")
		fmt.Scanln(&userId)
		fmt.Println("请输入用户密码")
		fmt.Scanln(&userPwd)

		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}

	} else if choice == 2 {
		fmt.Println("注册成功")
	}
}

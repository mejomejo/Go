package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

type Person struct {
	Name    string      `json:"name"`
	Age     int         `json:"age"`
	Address interface{} `json:"address"`
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			Street: "123 Main St.",
			City:   "Anytown",
			State:  "CA",
		},
	}

	// 将 Person 结构体序列化为 JSON 字符串
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	json.Unmarshal(data, &p)
	fmt.Println(p)
}

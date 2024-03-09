package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed user.json
var userList string

// User 结构体，与JSON对象的结构相对应
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// 定义一个User类型的切片
	var users []User

	// 解析json数据到结构体切片中
	err := json.Unmarshal([]byte(userList), &users)
	if err != nil {
		fmt.Println("Error parsing json data:", err)
	}

	// 遍历切片并打印每个用户的username和password
	for _, user := range users {
		fmt.Printf("Username: %s, Password: %s\n", user.Username, user.Password)
	}
}

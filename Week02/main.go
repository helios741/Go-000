package main

import (
	"Week02/utils"
	"Week02/service"
	"fmt"
)

func main() {
	var id = 23
	user, err := service.GetUserById(id)
	if utils.SqlIsNotFount(err) {
		// 处理为404或者空
		fmt.Println(user.Name, "Not found")
		fmt.Printf("%+v", err)
		return
	} else if err != nil {
		// 处理为500或者参数错误
		fmt.Println(user.Name, "find error")
		fmt.Printf("%+v", err)
		return
	}
	// 处理user逻辑
	fmt.Println(user)
}


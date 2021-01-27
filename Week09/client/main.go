package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	input := bufio.NewReader(os.Stdin)

	for {
		in, _ := input.ReadString('\n')
		_, err = conn.Write([]byte(in)) // 发送数据
		fmt.Println("ininiini: ", in)
		if err != nil {
			fmt.Println("发送失败")
			continue
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

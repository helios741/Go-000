package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)


func readProcess(conn net.Conn, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}
		fmt.Println("[readProcess] line: ", string(line))
		ch <- string(line)
	}
}

func writeProcess(conn net.Conn, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()
	writer := bufio.NewWriter(conn)
	for {
		line := <- ch
		fmt.Println("[writeProcess] line: ", line)
		if _, err := writer.WriteString(fmt.Sprintf("hello: %s", line)); err != nil {
				fmt.Println("write error: ", err)
				break
		}
		if err := writer.Flush(); err != nil {
			fmt.Println("write flush error: ", err)
			break
		}
	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
		return
	}
	// 等待信号之后关闭监听
	q := make(chan os.Signal)
	go func() {
		signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
		<- q

		if err = listen.Close(); err != nil {
			fmt.Println("listen clost error: ", err)
		}
	}()

	// 为的是等待所有的连接都退出之后再退出
	var wg = sync.WaitGroup{}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn error: ", err)
			break
		}
		ch := make(chan string)
		wg.Add(2)
		go readProcess(conn, ch, &wg)
		go writeProcess(conn, ch, &wg)
	}

	wg.Wait()

	fmt.Println("程序退出")
}

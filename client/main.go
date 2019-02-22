package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func doTask(conn net.Conn) {
	for {
		fmt.Fprintf(conn, "测试连接建立是否成功")
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("获取信息错误")
			break
		} else {
			fmt.Println("获取信息： ", msg)
		}
		time.Sleep(3 * time.Second)
	}
}
func main() {
	hostInfo := "127.0.0.1:10000"
	for {
		conn, err := net.Dial("tcp", hostInfo)
		fmt.Print("连接 (", hostInfo)
		if err != nil {
			fmt.Println(") 错误")
		} else {
			fmt.Println(") 成功")
			defer conn.Close()
			doTask(conn)
		}
		time.Sleep(3 * time.Second)
	}
}

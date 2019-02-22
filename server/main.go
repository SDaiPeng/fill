package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	bindInfo := ":10000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", bindInfo)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Println("等待连接")
	for {
		cc, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(cc)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(3 * time.Minute))
	request := make([]byte, 1024)
	defer conn.Close()
	for {
		recvLen, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if recvLen == 0 {
			break
		}
		recvData := strings.TrimSpace(string(request[:recvLen]))
		fmt.Println("获取长度 : ", recvLen)
		fmt.Println("获取数据 : " + recvData)
		conn.Write([]byte("测试一下BUG\n"))
		request = make([]byte, 1024)
	}
}

package main

import (
	"github.com/kataras/iris"
	"net"
)

const(
	ip = ""
	port = 3333
)

func main() {
	app := iris.New()

	l, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip),port,""})
	if err != nil {
		panic(err)
	}
	app.Run(iris.Listener(l))
}

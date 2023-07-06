package main

import (
	"log"
	"net"

	"module35.8.1/pkg/api"
)

var (
	proto = "tcp"
	port  = ":8080"
)

func main() {

	listener, err := net.Listen(proto, port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("new connect")
		go api.HandleConn(conn)
	}
}

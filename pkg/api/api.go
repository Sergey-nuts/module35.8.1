package api

import (
	"bufio"
	"log"
	"net"
	"strings"

	"module35.8.1/pkg/proverb"
)

// HandleConn обработчик соединения
// "exit" to close connection
func HandleConn(conn net.Conn) {
	defer conn.Close()
	done := make(chan struct{})
	reader := bufio.NewReader(conn)
	go proverb.RandProverb(done, conn)
	for {
		b, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}

		msg := strings.TrimSuffix(string(b), "\n")
		msg = strings.TrimSuffix(msg, "\r")
		log.Println("new message", msg)
		if strings.EqualFold(msg, "exit") {
			log.Println("receieve message `exit`")
			close(done)
			return
		}
	}
}

package echo

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct{}

func (self *Server) Listen(port int) {
	conn, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	check(err)

	var clients []net.Conn
	input := make(chan []byte, 10)

	go func() { // Chat Manager
		for {
			message := <-input
			for _, client := range clients {
				client.Write(message)
			}
		}
	}()

	for {
		client, err := conn.Accept()
		check(err)

		clients = append(clients, client)
		go handleClient(client, input)
	}
}

func handleClient(client net.Conn, input chan []byte) {
	for {
		buf := make([]byte, 4096)
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}
		// client.Write(buf) // single user only
		input <- buf // chat service
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

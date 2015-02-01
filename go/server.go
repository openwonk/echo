package echo

import (
	"fmt"
	"net"
	"strconv"
)

// Set socket parameters
// Create socket object
// Bind socket to port
// Listen and accept connections

type Server struct {
	// Host string
	Port int
}

func (s *Server) Listen() {
	conn, err := net.Listen("tcp", ":"+strconv.Itoa(s.Port))
	if err != nil {
		fmt.Println(err)
	}

	var clients []net.Conn
	input := make(chan []byte, 10)

	for {
		client, err := conn.Accept()
		if err != nil {
			continue
		}

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
		client.Write(buf)
	}
}

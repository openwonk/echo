package echo

import (
	// "bytes"
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	// users  = []int{}
	cmds   = make(map[int]string)
	store  = make(map[string]string)
	backup = make(map[string]string)
)

type Server struct{}

func (self *Server) Listen(port int) {
	conn, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	check(err)

	// var clients []net.Conn
	clients := make(map[string]net.Conn)
	input := make(chan []byte, 10)

	go func() { // CLIENT TERMINAL
		for {
			// PROMPT
			for _, client := range clients {
				client.Write([]byte(">> "))
			}

			// RECEIVE MESSAGES
			message := string(<-input)
			r, v := regexp.MustCompile("[^a-z|0-9|A-Z| ]"), ""
			message = r.ReplaceAllString(message, "")

			inputArray := strings.Split(message, " ")

			if len(inputArray) > 1 {
				// 	v = r.ReplaceAllString(inputArray[1], "")
				v = inputArray[1]

			}

			fmt.Println("'" + message + "'")
			fmt.Println("'" + inputArray[0] + "'")
			fmt.Println(len(inputArray[0]))

			// PARSE MESSAGES
			switch strings.ToLower(inputArray[0]) {
			case "set":
				if len(inputArray) < 3 {
					message = "set error"
				} else {
					store[inputArray[1]] = inputArray[2]
					message = inputArray[1] + " = " + inputArray[2]
				}

			case "get":
				message = store[v]
			case "unset":
				delete(store, v)
				message = ""
			case "all":
				message = ""
				for key, val := range store {
					message += key + ":" + val + "\n"
				}

			case "cmds":
				for k, v := range cmds {
					fmt.Println(k, v)
				}

			default:
				message = "entry error"

			}

			message = message + "\n"
			for _, client := range clients {
				client.Write([]byte(message))
				// fmt.Println(store)
			}
		}
	}()

	go func() { // Server Terminal
		fmt.Println("Listing on :" + strconv.Itoa(port))
		for {

			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("$ ")
			scanner.Scan()

			text := scanner.Text()
			r := regexp.MustCompile("[^a-z|0-9|A-Z| ]") // recalibrate to include "?!"
			text = r.ReplaceAllString(text, "")

			textArray := strings.Split(text, " ")

			switch strings.ToLower(textArray[0]) {
			case "quit", "exit":
				os.Exit(1)
			case "all":
				fmt.Println(store)
			case "users":
				for k, v := range clients {
					fmt.Println(k, v)
				}
			case "kill":
				fmt.Println("killing signal @ ", "'"+textArray[1]+"'")
				clients[textArray[1]].Write([]byte("\nAdmin closing connection... Goodbye\n"))
				clients[textArray[1]].Close()
				delete(clients, textArray[1])
			case "write":
				message := "\nAdmin => " + strings.Join(textArray[2:], " ") + "\n>> "
				clients[textArray[1]].Write([]byte(message))
			case "restart":

			default:
				fmt.Println("input error")
			}

		}
	}()

	for {

		client, err := conn.Accept()
		check(err)

		// client_port, _ := strconv.Atoi(client.RemoteAddr().String()[6:])
		// users = append(users, client_port)

		client.Write([]byte(">> "))

		// clients = append(clients, client)
		clients[client.RemoteAddr().String()[6:]] = client
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

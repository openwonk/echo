package echo

import (
	// "bytes"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	cmds   = make(map[int]string)
	store  = make(map[string]string)
	backup = make(map[string]string)
)

type Server struct{}

func (self *Server) Listen(port int) {
	conn, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	check(err)

	var clients []net.Conn
	input := make(chan []byte, 10)

	go func() { // Chat Manager
		for {
			for _, client := range clients {
				client.Write([]byte(">> "))
			}

			message := string(<-input)
			inputArray := strings.Split(message, " ")
			r := regexp.MustCompile("[^a-z|0-9|A-Z| ]")

			fmt.Println(len(message), message)

			switch v := r.ReplaceAllString(inputArray[1], ""); strings.ToLower(inputArray[0]) {
			case "set":
				store[inputArray[1]] = inputArray[2]
			case "get":
				message = store[v]

			case "unset":
				delete(store, v)

			case "all":
				for k, v := range store {
					fmt.Println(k, v)
				}

			case "cmds":
				for k, v := range cmds {
					fmt.Println(k, v)
				}

			case "equals":
				c := 0
				for _, v := range store {
					if v == inputArray[1] {
						c++
					}
				}
				fmt.Println(c)

			case "greater":
				c := 0
				i, err := strconv.Atoi(inputArray[1])
				check(err)

				for _, v := range store {
					n, _ := strconv.Atoi(v)
					if n > i {
						c++
					}
				}

				fmt.Println(c)

			case "lesser":
				c := 0
				i, err := strconv.Atoi(inputArray[1])
				check(err)

				for _, v := range store {
					n, _ := strconv.Atoi(v)
					if n < i {
						c++
					}
				}

				fmt.Println(c)

			case "add":
				i, err := strconv.Atoi(inputArray[1])
				if err != nil {
					i, _ = strconv.Atoi(store[inputArray[1]])
				}

				j, err := strconv.Atoi(inputArray[2])
				if err != nil {
					j, _ = strconv.Atoi(store[inputArray[2]])
				}

				m := i + j
				fmt.Println(inputArray[1], "+", inputArray[2], "=", m)

			case "subtract":
				i, err := strconv.Atoi(inputArray[1])
				if err != nil {
					i, _ = strconv.Atoi(store[inputArray[1]])
				}

				j, err := strconv.Atoi(inputArray[2])
				if err != nil {
					j, _ = strconv.Atoi(store[inputArray[2]])
				}

				m := i - j
				fmt.Println(inputArray[1], "-", inputArray[2], "=", m)

			case "clear":
				cmds = make(map[int]string)
				store = make(map[string]string)

			// case "begin", "commit":
			// 	copy(&backup, store)

			// case "restore":
			// 	copy(&store, backup)

			case "save":
				filename := inputArray[1]
				// err = ioutil.WriteFile(filename, store)
				f, err := os.Create(filename)
				check(err)
				defer f.Close()

				// TODO: convert store to JSON
				// then save to local file
				sample := []byte{115, 111, 109, 101, 10}

				f.Write(sample)
				fmt.Printf("Saved to %s\n", filename)

			case "exit":
				fmt.Println("Goodbye\n")
				os.Exit(0)

				// // only use in devmode
				// case "getbackup":
				// 	fmt.Println(backup)

			}

			for _, client := range clients {
				client.Write([]byte(message))
				// fmt.Println(store)
			}
		}
	}()

	for {

		client, err := conn.Accept()
		check(err)

		client.Write([]byte(">> "))

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

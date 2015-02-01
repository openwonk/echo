package echo

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

// Set socket parameters
// Create socket object

type Client struct{}

func (c *Client) Connect(port int) {
	conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(port))
	check(err)
	defer conn.Close()

	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	for {
		go func() {
			for {
				buf := make([]byte, 4096)
				numbytes, err := conn.Read(buf)
				if numbytes == 0 || err != nil {
					return
				}

				var s string
				for _, v := range buf {
					s += string(v)
				}

				fmt.Println(s)
			}
		}()

		scanner := bufio.NewScanner(os.Stdin)
		// fmt.Print("> ")
		scanner.Scan()
		outgoing := []byte(scanner.Text() + "\n")
		conn.Write(outgoing)
	}

}

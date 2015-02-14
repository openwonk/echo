package main

import (
	echo "./go"
)

func main() {
	server := echo.Server{}
	server.Listen(8080)
}

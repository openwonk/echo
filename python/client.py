#!/usr/bin/python

import socket, sys

HOST = "127.0.0.1"
PORT = 8080

client = socket.socket(
	socket.AF_INET,
	socket.SOCK_STREAM,
	socket.IPPROTO_IP)

print "connecting to :" + str(PORT)
client.connect((HOST, PORT))

while True:
	outgoing = raw_input("> ")
	client.sendall(outgoing)
	if outgoing == u"exit" or outgoing == u"quit":
		client.close()
		break
	print client.recv(1000)



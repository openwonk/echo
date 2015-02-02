#!/usr/bin/python

import socket, sys

def client(msg, log_buffer=sys.stderr):
	port = sys.argv[2] if len(sys.argv) > 2 else 8080

	client = socket.socket(
		socket.AF_INET,
		socket.SOCK_STREAM,
		socket.IPPROTO_IP)

	print "connecting to :" + str(port)
	client.connect(('127.0.0.1', port))

	while True:
		outgoing = raw_input("> ")
		client.sendall(outgoing)
		if outgoing == u"exit" or outgoing == u"quit":
			client.close()
			break
		print client.recv(1000)



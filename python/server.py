#!/usr/bin/python

import socket, sys

port = sys.argv[2] if len(sys.argv) > 2 else 8080

server = socket.socket(
	socket.AF_INET,
	socket.SOCK_STREAM,
	socket.IPPROTO_IP)

server.bind(('127.0.0.1', port))
print "listening on :" + str(port)
server.listen(4)

def Handler(conn, addr):
	client = str(addr[1])
	print "client:" + client + " connected"
	broken_pipe = 0
	session = True

	while session:
		
		incoming = conn.recv(1000)
		exit = incoming.replace("\n","")

		if incoming.replace("\n","") == "":
			broken_pipe += 1
			if broken_pipe > 5:
				session = False
		elif exit == u"exit" or exit == u"quit":
			session = False
		else:
			broken_pipe = 0
			print "client:" + client + " says '" + incoming + "'"
			conn.sendall(incoming)

	conn.close()
	print "client:" + client + " disconnected"
	return


while True:
	conn, addr = server.accept()
	Handler(conn, addr)



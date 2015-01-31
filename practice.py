#!/usr/bin/python

import socket

TARGET_SITE = 'crisewing.com'

streams = [info
	for info in socket.getaddrinfo(TARGET_SITE, 'http')
	if info[1] == socket.SOCK_STREAM]

# prepare request
info = streams[0]
new_socket = socket.socket(*info[:3])
new_socket.connect(info[-1])
msg = "GET / HTTP/1.1\r\n"
msg += "Host: " + TARGET_SITE + "\r\n\r\n"
# send request
new_socket.sendall(msg)


# prepare for extracting response
buffsize = 4096
response = ""
done = False
# loop through and concat site
while not done:
	msg_part = new_socket.recv(buffsize)
	# if msg_part: # should that be zeroed out?
	if len(msg_part) < buffsize: # should that be zeroed out?
		done = True
		new_socket.close()
	response += msg_part

print len(response)


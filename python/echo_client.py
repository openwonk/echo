#!/usr/bin/python

import socket, sys

HOST = "127.0.0.1"
PORT = 8080

def client(msg, log_buffer=sys.stderr):
    server_address = (HOST, PORT)
    sock = socket.socket(
        socket.AF_INET,
        socket.SOCK_STREAM,
        socket.IPPROTO_IP)
    # print 'connecting to {0} port {1}'.format(*server_address)
    print >>log_buffer, 'connecting to {0} port {1}'.format(*server_address)
    sock.connect((HOST, PORT))

    try:
        print 'sending "{0}"'.format(msg)  
        print >>log_buffer, 'sending "{0}"'.format(msg)
        sock.sendall(msg)
        done = False
        response = ""

        while not done:
            chunk = sock.recv(16)
            # print chunk
            # print 'received({0}) "{1}"'.format(len(chunk), chunk)
            print >>log_buffer, 'received "{0}"'.format(chunk)
            if len(chunk) >= 16 :
                response += chunk
            else:
                response += chunk.strip()
                done = True
                break
        print "Full message: ", response

    finally:
        print >>log_buffer, 'closing socket'
        sock.close()

if __name__ == '__main__':
    if len(sys.argv) != 2:
        usg = '\nusage: python echo_client.py "this is my message"\n'
        print >>sys.stderr, usg
        sys.exit(1)

    msg = sys.argv[1]
    client(msg)
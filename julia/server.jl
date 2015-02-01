function +(a::String, b::String)
	a * b
end

port = 8080
server = listen(port)
println("listen @ " + string(port))

while true
	conn = accept(server)
	@async begin
		try
			while true
				incoming = readline(conn)
				print("> ", incoming)
				write(conn, incoming)
			end
		catch err
			print("error: $err")
		end

	end
end


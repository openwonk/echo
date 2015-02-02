function +(a::String, b::String)
	a * b
end

function Echo(port)
	# port = 8080
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
					# incoming = replace(input, "\n", "")
					# if incoming == ":exit" || input == ":quit"
				end
			catch err
				print("error: $err")
			end

		end
	end
end

Echo(8080)


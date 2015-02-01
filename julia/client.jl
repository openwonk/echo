client = connect(8080)

while true
	print("> ")
	input = readline(STDIN)
	input = replace(input, "\n", "")
	println(client, input)
	if input == ":exit" || input == ":quit"
		close(client)
		exit(1)
	else
		write(STDOUT,readline(client))
	end
end

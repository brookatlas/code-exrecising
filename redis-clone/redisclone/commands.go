package redisclone

func command(command_array []string) []byte {

	var initial_command string

	if len(command_array) == 0 {
		initial_command = "DOCS"
	} else {
		initial_command = command_array[0]
	}

	switch initial_command {
	case "DOCS":
		return command_docs(command_array)
	default:
		return command_docs(command_array)
	}
}

func command_docs(command_array []string) []byte {
	return []byte("*2\r\n$3\r\nSET\r\n%1\r\n$7\r\nsummary\r\n$4\r\nthis\r\n")
}

func ping() []byte {
	response := writeSimpleString("PONG")

	return response
}

package redisclone

type RedisCloneCommandInfo struct {
	CommandName    string
	CommandSummary string
}

var SET_COMMAND_INFO RedisCloneCommandInfo = RedisCloneCommandInfo{
	CommandName:    "SET",
	CommandSummary: "this is the set command summary",
}

func (c RedisCloneCommandInfo) getCommandDocs() []byte {
	summary_map := writeStringMap(map[string]string{
		"summary": c.CommandSummary,
	})
	command_string := writeBulkString(c.CommandName)
	formatted_response_string_array := []string{
		string(command_string),
		string(summary_map),
	}
	response_array := writeRawArray(formatted_response_string_array)

	return response_array
}

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
	response := SET_COMMAND_INFO.getCommandDocs()

	return response
}

func ping() []byte {
	response := writeSimpleString("PONG")

	return response
}

package redisclone

import (
	"fmt"
	"log"
)

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

func set(store *RedisCloneStore, command_array []string) []byte {
	if len(command_array) < 3 {
		return command_docs([]string{
			"COMMAND",
			"DOCS",
			"SET",
		})
	}

	key, value := command_array[1], command_array[2]

	response_ok := store.StoreSet(key, value)

	if !response_ok {
		error_message_format := "error while setting a key called: %s"
		error_message := fmt.Sprintf(error_message_format, key)

		log.Fatal(error_message)
		panic(1)
	}

	response := writeSimpleString("OK")

	return response
}

func get(store *RedisCloneStore, command_array []string) []byte {
	if len(command_array) < 2 {
		return command_docs([]string{
			"COMMAND",
			"DOCS",
			"GET",
		})
	}

	key := command_array[1]

	value := store.StoreGet(key)

	response := writeSimpleString(value)

	return response
}

func config(config *RedisCloneConfig, command_array []string) []byte {
	if len(command_array) < 3 {
		response := writeError("missing arguments")
		return response
	}

	sub_command := command_array[1]
	switch sub_command {
	case "GET":
		response := config_get(config, command_array)
		return response
	}

	response := writeBulkString("unknown sub command")

	return response
}

func config_get(config *RedisCloneConfig, command_array []string) []byte {
	config_name := command_array[2]
	switch config_name {
	case "save":
		response := writeRawArray(config.save)
		return response
	case "appendonly":
		response := writeSimpleString(config.appendonly)
		return response
	}

	response := writeError("unknown")

	return response
}

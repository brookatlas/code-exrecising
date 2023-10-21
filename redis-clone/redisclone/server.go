package redisclone

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

type RedisCloneClient struct {
	Host string
	Port int
}

func (c RedisCloneClient) Run() {
	connection_type := "tcp"

	connection_string := fmt.Sprintf(
		"%s:%d",
		c.Host,
		c.Port,
	)
	listener, err := net.Listen(
		connection_type,
		connection_string,
	)

	if err != nil {
		error_message := "there was a problem setting up a tcp listener over 6379. exiting ..."
		log.Fatal(
			error_message,
		)

		os.Exit(-1)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			error_message_format := "there was an error accepting an incoming tcp connection: %s"
			error_message := fmt.Sprintf(error_message_format, err)
			log.Fatal(error_message)

			panic(err)
		}

		go handleRequest(connection)
	}
}

func handleRequest(connection net.Conn) {

	for {
		buffer := make([]byte, 1024)

		_, err := connection.Read(buffer)
		if err != nil {
			error_message_format := "error while reading bytes from connection: %s"
			error_message := fmt.Sprintf(error_message_format, err)
			log.Fatal(error_message)
			panic(error_message)
		}

		command_array, _ := readValue(buffer)
		initial_command := command_array[0]
		var response_pointer *[]byte

		switch initial_command {
		case "PING":
			response := ping()
			response_pointer = &response
		case "COMMAND":
			response := command(command_array[1:])
			response_pointer = &response
		}

		if response_pointer != nil {
			connection.Write(*response_pointer)
		}
	}
}

func readValue(buffer []byte) ([]string, int) {
	var value_pointer *string
	var length_pointer *int
	var array_pointer *[]string

	first_byte := buffer[0]
	buffer_to_send := buffer[1:]

	switch first_byte {
	case '+':
		value, length := readSimpleString(buffer_to_send)
		value_pointer = &value
		length_pointer = &length
	case '-':
		value, length := readError(buffer_to_send)
		value_pointer = &value
		length_pointer = &length
	case ':':
		value, length := readInteger(buffer_to_send)
		value_pointer = &value
		length_pointer = &length
	case '$':
		value, length := readBulkString(buffer_to_send)
		value_pointer = &value
		length_pointer = &length
	case '*':
		value, length := readArray(buffer_to_send)
		array_pointer = &value
		length_pointer = &length
	default:
		error_message := "message started with an invalid format."
		log.Fatal(error_message)
		panic(1)
	}

	if value_pointer != nil {
		string_array := make([]string, 1)
		string_array[0] = *value_pointer
		array_pointer = &string_array
	}

	return *array_pointer, *length_pointer
}

func readSimpleString(buffer []byte) (string, int) {
	stop_index := bytes.Index(buffer, []byte("\r\n"))
	value := buffer[0:stop_index]

	return string(value), stop_index + 2
}

func readError(buffer []byte) (string, int) {
	stop_index := bytes.Index(buffer, []byte("\r\n"))
	value := buffer[0:stop_index]

	return string(value), stop_index + 2
}

func readInteger(buffer []byte) (string, int) {
	stop_index := bytes.Index(buffer, []byte("\r\n"))
	parsed_value := string(buffer[0:stop_index])

	return parsed_value, stop_index + 2
}

func readBulkString(buffer []byte) (string, int) {
	string_length, string_start_index := readInteger(buffer)

	parsed_string_length, _ := strconv.Atoi(string_length)

	string_raw := buffer[string_start_index : string_start_index+int(parsed_string_length)]

	return string(string_raw), string_start_index + int(parsed_string_length) + 3
}

func readArray(buffer []byte) ([]string, int) {
	buffer_offset := 0
	token_array := make([]string, 0)
	array_length, next_buffer_offset := readInteger(buffer[buffer_offset:])
	buffer_offset = next_buffer_offset

	array_length_parsed, _ := strconv.Atoi(array_length)

	for i := 0; i < array_length_parsed; i++ {
		current_value, next_buffer_offset := readValue(buffer[buffer_offset:])
		buffer_offset = buffer_offset + next_buffer_offset
		token_array = append(token_array, current_value[0])
	}

	return token_array, len(token_array)
}

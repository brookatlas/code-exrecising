package redisclone

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

type RedisCloneClient struct {
	Host   string
	Port   int
	Store  RedisCloneStore
	Config RedisCloneConfig
}

type RedisCloneConfig struct {
	save       []string
	appendonly string
}

func NewRedisCloneClient(host string, port int) *RedisCloneClient {
	redis_clone_client := RedisCloneClient{
		Host: host,
		Port: port,
		Store: RedisCloneStore{
			mu:   sync.RWMutex{},
			dict: map[string]string{},
		},
		Config: RedisCloneConfig{
			save: []string{
				"60",
				"1000",
			},
			appendonly: "no",
		},
	}

	return &redis_clone_client
}

func (c *RedisCloneClient) Run() {
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

		go c.handleRequest(connection)
	}
}

func (c *RedisCloneClient) handleRequest(connection net.Conn) {

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
		case "SET":
			response := set(&c.Store, command_array)
			response_pointer = &response
		case "GET":
			response := get(&c.Store, command_array)
			response_pointer = &response
		case "PING":
			response := ping()
			response_pointer = &response
		case "COMMAND":
			response := command(command_array[1:])
			response_pointer = &response
		case "CONFIG":
			response := config(&c.Config, command_array)
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

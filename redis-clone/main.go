package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

// use this for help: https://github.com/ngilles/codecrafters-redis-py/blob/master/app/main.py
// https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go
// https://codingchallenges.fyi/challenges/challenge-redis/

func main() {
	connection_type := "tcp"
	connection_port := 6379
	connection_host := "localhost"

	connection_string := fmt.Sprintf(
		"%s:%d",
		connection_host,
		connection_port,
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
	buffer := make([]byte, 1024)
	for {
		_, err := connection.Read(buffer)
		if err != nil {
			error_message_format := "error while reading bytes from connection: %s"
			error_message := fmt.Sprintf(error_message_format, err)
			log.Fatal(error_message)
			panic(error_message)
		}
	}

	// connection.Write(BYTE ARRAY HERE)
	// connection.Close()
}

func readRedisRequest(buffer []byte) []byte {
	first_byte := buffer[0]
	buffer_to_send := buffer[1:]
	if first_byte == '+' {
		return readSimpleString(buffer_to_send)
	}
	if first_byte == '-' {
		return readError(buffer_to_send)
	}
	if first_byte == ':' {
		return readInteger(buffer_to_send)
	}
	if first_byte == '$' {
		return readBulkString(buffer_to_send)
	}
	if first_byte == '*' {
		return readArray(buffer_to_send)
	}
}

func readSimpleString(buffer []byte) []byte {
	stop_index := bytes.Index(buffer, []byte{"\r", "\n"})
	value := buffer[0:stop_index]
	length := len(buffer[0:stop_index]) + 2

	return value, length
}

func readError(buffer []byte) []byte {

}

func readInteger(buffer []byte) []byte {

}

func readBulkString(buffer []byte) []byte {

}

func readArray(buffer []byte) []byte {

}

package redisclone

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func readSimpleString(buffer []byte) (string, int) {
	stop_index := bytes.Index(buffer, []byte("\r\n"))
	value := buffer[0:stop_index]

	return string(value), stop_index + 2
}

func writeSimpleString(str string) []byte {
	format := "+%s\r\n"
	formatted_string := fmt.Sprintf(format, str)

	return []byte(formatted_string)
}

func readError(buffer []byte) (string, int) {
	stop_index := bytes.Index(buffer, []byte("\r\n"))
	value := buffer[0:stop_index]

	return string(value), stop_index + 2
}

func writeError(str string) []byte {
	format := "-%s\r\n"
	formatted_string := fmt.Sprintf(format, str)

	return []byte(formatted_string)
}

func readInteger(buffer []byte) (string, int) {
	stop_index := bytes.Index(buffer, []byte("\r\n"))
	parsed_value := string(buffer[0:stop_index])

	return parsed_value, stop_index + 2
}

func writeInteger(num int) []byte {
	format := ":%d\r\n"
	formatted_string := fmt.Sprintf(format, num)

	return []byte(formatted_string)
}

func readBulkString(buffer []byte) (string, int) {
	string_length, string_start_index := readInteger(buffer)

	parsed_string_length, _ := strconv.Atoi(string_length)

	string_raw := buffer[string_start_index : string_start_index+int(parsed_string_length)]

	return string(string_raw), string_start_index + int(parsed_string_length) + 3
}

func writeBulkString(str string) []byte {
	format := "$%d\r\n%s\r\n"
	formatted_string := fmt.Sprintf(format, len(str), str)

	return []byte(formatted_string)
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

func writeRawArray(string_array []string) []byte {
	var builder strings.Builder
	array_size_format := "*%d\r\n"
	array_size_string := fmt.Sprintf(array_size_format, len(string_array))
	builder.Write([]byte(array_size_string))
	for _, str := range string_array {
		builder.Write([]byte(str))
	}

	return []byte(builder.String())
}

func writeStringMap(mp map[string]string) []byte {
	var builder strings.Builder
	map_size_format := "%%%d\r\n"
	map_size_string := fmt.Sprintf(map_size_format, len(mp))
	builder.Write([]byte(map_size_string))
	for key, value := range mp {
		current_key_bytes := writeBulkString(key)
		builder.Write(current_key_bytes)
		current_value_bytes := writeBulkString(value)
		builder.Write(current_value_bytes)
	}

	return []byte(builder.String())
}

package main

import (
	"fmt"
)

func get_status_string(status string) string {
	return "+" + status + "\r\n"
}

func get_error_string(error string) string {
	return "-" + error + "\r\n"
}

func get_number_string(n int) string {
	return fmt.Sprintf(":%d\r\n", n)
}

func get_empty_bulk_string() string {
	return "$-1\r\n"
}

func get_bulk_string(content string) string {
	n := fmt.Sprintf("$%d\r\n", len(content))
	return n + content + "\r\n"
}

func get_empty_bulks_string() string {
	return "*0\r\n"
}

func get_bulks_string(content []string) string {
	if len(content) == 0 {
		return get_empty_bulks_string()
	}
	result := fmt.Sprintf("*%d\r\n", len(content))
	for i := 0; i < len(content); i++ {
		result = result + get_bulk_string(content[i])
	}
	return result
}

func get_request_string(cmd string, args []string) string {
	result := fmt.Sprintf("*%d\r\n", len(args)+1)
	result = result + get_bulk_string(cmd)
	for i := 0; i < len(args); i++ {
		result = result + get_bulk_string(args[i])
	}
	return result
}

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

func get_empty_bulk_string() string {
	return "$-1\r\n"
}

func get_bulk_string(content string) string {
	n := fmt.Sprintf("$%d\r\n", len(content))
	return n + content + "\r\n"
}

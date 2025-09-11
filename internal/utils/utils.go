package utils

import (
	"fmt"
	"strings"
)

func EncodeCommand(parts ...string) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("*%d\r\n", len(parts)))

	for _, p := range parts {
		sb.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(p), p))
	}

	return sb.String()
}

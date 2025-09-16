package netutil

import (
	"fmt"
	"strconv"
	"strings"
)

// ParsePort converts a string to a valid TCP/UDP port number.
// It trims spaces and returns an error if the port is out of range.
func ParsePort(s string) (int, error) {
	p, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil || p < 1 || p > 65535 {
		return 0, fmt.Errorf("invalid port")
	}
	return p, nil
}

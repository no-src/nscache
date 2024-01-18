package unit

import "github.com/no-src/nsgo/unit"

// ParseBytes parse the string to bytes
func ParseBytes(s string) (int, error) {
	bytes, _, err := unit.ParseBytes(s)
	if err != nil {
		return 0, err
	}
	return int(bytes), nil
}

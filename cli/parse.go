package cli

import "strings"

func getCommand(name string) *command {
	return commands[strings.ToLower(name)]
}

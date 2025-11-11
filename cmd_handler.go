package main

import "strings"

type ReddishCmdHandler struct{}

var optype = map[string]bool{
	"+": true,
	"-": true,
	":": true,
	"$": true,
	"*": true,
	"_": true,
	"#": true,
	",": true,
	"(": true,
	"!": true,
	"=": true,
	"%": true,
	"|": true,
	"~": true,
	">": true,
}

func (ch *ReddishCmdHandler) ParseArgs(input string) ([]string, error) {
	tokens := make([]string, 0)
	if len(input) == 0 {
		return tokens, nil
	}

	parts := strings.Split(input, "\r\n")
	parts = parts[:len(parts)-1]

	for _, token := range parts {
		if !optype[string(token[0])] {
			tokens = append(tokens, token)
		}
	}

	return tokens, nil
}

func (ch *ReddishCmdHandler) HandleCommand(cmdTokens []string) (string, error) {
	cmd := strings.ToUpper(cmdTokens[0])

	// Initialise DB

	switch cmd {
	case "PING":
		return ch.handlePing(cmdTokens)
	default:
		return "-Error: Unknown command\r\n", nil
	}
}

func (ch *ReddishCmdHandler) handlePing(tokens []string) (string, error) {
	if len(tokens) > 2 {
		return "-Error: PING either takes no or 1 arg\r\n", nil
	}
	if len(tokens) == 1 {
		return "+PONG\r\n", nil
	}
	return "+" + tokens[1] + "\r\n", nil
}

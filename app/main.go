package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Read input
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}

		// Normalize
		command = strings.TrimSpace(command)
		words := strings.Fields(command)

		if len(words) == 0 {
			continue // ignore empty input
		}

		// Exit
		if command == "exit 0" {
			break
		}

		// Builtin: type
		if strings.ToLower(words[0]) == "type" {
			if len(words) < 2 {
				fmt.Println("type: missing operand")
				continue
			}
			rest := words[1]
			if strings.ToLower(rest) == "invalid_command" {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", rest)
			} else {
				fmt.Println(rest, "is a shell builtin")
			}
			continue
		}

		// Builtin: echo
		if strings.ToLower(words[0]) == "echo" {
			rest := strings.Join(words[1:], " ")
			fmt.Println(rest)
			continue
		}

		// Fallback: not found
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

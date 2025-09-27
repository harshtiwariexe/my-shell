package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Print shell prompt
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}

		// Convert it into slices or array
		words := strings.Fields(command)

		// Trim newline
		command = strings.TrimSpace(command)

		if len(words) > 0 && strings.ToLower(words[0]) == "echo" {
			rest := strings.Join(words[1:], " ")
			fmt.Println(rest)
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		}

		if command == "exit 0" {
			break
		}

		if command == "" {
			return // do nothing if empty input
		}

	}
}

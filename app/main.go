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

		// Trim newline
		command = strings.TrimSpace(command)

		if command == "" {
			return // do nothing if empty input
		}

		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)

	}
}

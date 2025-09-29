package main

import (
	"fmt"
	"strings"
)

func isBuiltin(cmd string) bool {
	switch strings.ToLower(cmd) {
	case "cd", "echo", "type", "exit", "pwd":
		return true
	default:
		return false
	}
}

func handleEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func handleType(cmd string) {

	if isBuiltin(cmd) {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return
	}

	if path := findExecutable(cmd); path != "" {
		fmt.Printf("%s is %s\n", cmd, path)
		return
	}

	fmt.Printf("%s: not found\n", cmd)
}

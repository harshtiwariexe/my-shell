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
		fmt.Fprintf(os.Stdout, "$ ")
		line, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		words := strings.Fields(strings.TrimSpace(line))

		if len(words) == 0 {
			continue
		}

		switch words[0] {
		case "exit":
			return

		case "echo":
			handleEcho(words[1:])

		case "type":
			if len(words) < 2 {
				fmt.Println("type: missing operands")
				continue
			}
			handleType(words[1])
			continue

		default:
			fmt.Printf("%s: not found\n", words[0])
		}
	}
}

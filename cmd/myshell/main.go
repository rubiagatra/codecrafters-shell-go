package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	builtIns := []string{"echo", "type", "exit"}
	for {
		fmt.Fprint(os.Stdout, "$ ")
		message, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		message = strings.TrimSpace(message)
		commands := strings.Split(message, " ")
		switch commands[0] {
		case "exit":
			code, err := strconv.Atoi(commands[1])
			if err != nil {
				os.Exit(1)
			}
			os.Exit(code)
		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(commands[1:], " "))
		case "type":
			if slices.Contains(builtIns, commands[1]) {
				fmt.Fprintf(os.Stdout, "%s %s\n", strings.Join(commands[1:], " "), "is a shell builtin")
				continue
			}

			fmt.Fprintf(os.Stdout, "%s: not found\n", commands[1])
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", message)
		}
	}
}

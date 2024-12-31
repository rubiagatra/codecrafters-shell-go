package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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
			found := false
			switch commands[1] {
			case "exit", "echo", "type":
				fmt.Printf("%s is a shell builtin\n", commands[1])
			default:
				paths := strings.Split(os.Getenv("PATH"), ":")
				for _, path := range paths {
					entries, _ := os.ReadDir(path)
					for _, file := range entries {
						if file.Name() == commands[1] {
							fmt.Printf("%s is %s/%s\n", commands[1], path, commands[1])
							found = true
							break
						}
					}
					if found {
						break
					}
				}
				if !found {
					fmt.Fprintf(os.Stdout, "%s: not found\n", commands[1])
				}
			}

		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", message)
		}
	}
}

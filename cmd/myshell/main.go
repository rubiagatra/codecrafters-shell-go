package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
					exec := path + "/" + commands[1]
					if _, err := os.Stat(exec); err == nil {
						fmt.Fprintf(os.Stdout, "%v is %v\n", commands[1], exec)
						found = true
					}
				}
				if !found {
					fmt.Fprintf(os.Stdout, "%s: not found\n", commands[1])
				}
			}

		default:
			found := false
			paths := strings.Split(os.Getenv("PATH"), ":")
			for _, path := range paths {
				execPath := path + "/" + commands[0]
				cmd := exec.Command(execPath, commands[1:]...)
				out, _ := cmd.CombinedOutput()
				if err == nil && len(out) != 0 {
					fmt.Fprintf(os.Stdout, "%s", out)
					found = true
				}
			}

			if !found {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", message)
			}
		}
	}
}

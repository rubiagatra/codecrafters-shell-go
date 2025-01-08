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
		case "pwd":
			cwd, _ := os.Getwd()
			fmt.Fprintf(os.Stdout, "%s\n", cwd)
		case "cd":
			if len(commands) == 1 {
				commands = append(commands, "~")
			}

			switch commands[1] {
			case "~":
				home := os.Getenv("HOME")
				err = os.Chdir(home)
				if err != nil {
					fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", commands[1])
				}
			default:
				err = os.Chdir(commands[1])
				if err != nil {
					fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", commands[1])
				}

			}

		case "type":
			found := false
			switch commands[1] {
			case "exit", "echo", "type", "pwd", "cd":
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
			cmd := exec.Command(commands[0], commands[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", commands[0])
			}
		}
	}
}

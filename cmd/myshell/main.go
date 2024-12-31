package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	defaultCommands := []string{"echo"}
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
			os.Exit(1)
		}

		strippedCommand := strings.Split(command[:len(command)-1], " ")
		if strippedCommand[0] == "exit" && strippedCommand[1] == "0" {
			os.Exit(0)
		}

		if slices.Contains(defaultCommands, strippedCommand[0]) {
			fmt.Println(strings.Join(strippedCommand[1:], " "))
			continue
		}

		fmt.Println(strippedCommand[0] + ": command not found")
	}

}

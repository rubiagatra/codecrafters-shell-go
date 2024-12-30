package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	defaultCommands := []string{"ls"}
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
			os.Exit(1)
		}

		strippedCommand := command[:len(command)-1]
		if strippedCommand == "exit" {
			fmt.Println("Exiting")
			os.Exit(0)
		}

		if slices.Contains(defaultCommands, strippedCommand) {
			fmt.Println(strippedCommand)
			continue
		}

		fmt.Println(strippedCommand + ": command not found")
	}

}

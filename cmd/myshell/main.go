package main

import (
	"bufio"
	"fmt"
	"os"
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

		for _, v := range defaultCommands {
			if v == strippedCommand {
				fmt.Println("Do the Command")
				continue
			}
		}

		fmt.Println(strippedCommand + ": command not found")
	}

}

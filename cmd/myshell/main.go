package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprint(os.Stdout, "Error reading input\n")
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%s: command not found\n", input[:len(input)-1])
}

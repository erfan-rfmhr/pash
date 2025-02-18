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
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprint(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		full_command := strings.Fields(command)
		if len(full_command) == 0 {
			os.Exit(0)
		}
		parent := full_command[0]
		switch parent {
		case "exit":
			exitCode, _ := strconv.Atoi(strings.TrimSpace(full_command[1]))
			os.Exit(exitCode)
		case "echo":
			fmt.Println(full_command[1])
		default:
			fmt.Fprintln(os.Stdout, strings.TrimSpace(command)+": command not found")
		}
	}
}

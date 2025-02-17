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
		splited := strings.Split(command, " ")
		parent := splited[0]
		switch parent {
		case "\n":
			os.Exit(0)
		case "exit":
			exitCode, _ := strconv.Atoi(strings.TrimSpace(splited[1]))
			os.Exit(exitCode)
		}
		fmt.Fprintln(os.Stdout, command[:len(command)-1]+": command not found")
	}
}

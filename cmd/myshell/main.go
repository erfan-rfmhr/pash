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
			panic(err)
		}
		full_command := strings.Fields(command)
		if len(full_command) == 0 {
			os.Exit(0)
		}
		parent := full_command[0]
		switch parent {

		case "type":
			type_command(full_command[1])
			
		case "exit":
			exit(full_command[1])
			
		case "echo":
			echo(full_command[1:])

		default:
			fmt.Fprintln(os.Stdout, strings.TrimSpace(command)+": command not found")
		}
	}
}

func type_command(command string) {
	shell_built_in_commands := [...]string{"echo", "exit", "type"}
	for i := 0; i<len(shell_built_in_commands); i++ {
		if shell_built_in_commands[i] == command {
			fmt.Println(command + " is a shell builtin")
			return
		}
	}
	fmt.Println(command + ": not found")
}

func exit(code string) {
	exitCode, _ := strconv.Atoi(strings.TrimSpace(code))
	os.Exit(exitCode)
}

func echo(args []string) {
	n := len(args)
	for i := 0; i < n; i++ {
		if i < n-1 {
			fmt.Print(args[i] + " ")
		} else {
			fmt.Print(args[i])
		}
	}
	fmt.Println()
}
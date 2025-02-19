package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
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
		args := full_command[1:]
		
		switch parent {

		case "type":
			type_command(full_command[1])

		case "exit":
			exit(full_command[1])

		default:
			run(parent, args)
		}
	}
}

func run(command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(strings.TrimSpace(command)+": command not found")	
	}
}

func type_command(command string) {
	shell_built_in_commands := []string{"echo", "exit", "type"}
	if slices.Contains(shell_built_in_commands, command) {
		fmt.Println(command + " is a shell builtin")
		return
	} else {
		if path, err := exec.LookPath(command); err == nil {
			fmt.Println(command + " is " + path)
			return
		}
	}
	fmt.Println(command + ": not found")
}

func exit(code string) {
	exitCode, _ := strconv.Atoi(strings.TrimSpace(code))
	os.Exit(exitCode)
}

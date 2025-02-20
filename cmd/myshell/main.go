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
		fullCommand := strings.Fields(command)
		if len(fullCommand) == 0 {
			os.Exit(0)
		}
		parent := fullCommand[0]
		args := fullCommand[1:]

		switch parent {
		case "type":
			typeCommand(fullCommand[1])

		case "exit":
			exit(fullCommand[1])

		case "pwd":
			wd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(wd)
			}

		case "cd":
			if len(args) != 0 {
				cd(args[0])
			} else {
				cd("~")
			}

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
		fmt.Println(strings.TrimSpace(command) + ": command not found")
	}
}

func typeCommand(command string) {
	shellBuiltInCommands := []string{"echo", "exit", "type", "pwd"}
	if slices.Contains(shellBuiltInCommands, command) {
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

func cd(dir string) {
	switch dir {
	case "..":
		err := os.Chdir("../")
		if err != nil {
			fmt.Println(err)
		}

	case "~":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		}
		err = os.Chdir(homeDir)
		if err != nil {
			fmt.Println(err)
		}

	default:
		err := os.Chdir(dir)
		if err != nil {
			fmt.Println("cd: " + dir + ": No such file or directory")
		}
	}
	currDir, _ := os.Getwd()
	os.Setenv("PWD", currDir)
}

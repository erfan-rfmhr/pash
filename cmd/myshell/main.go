package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/term"
)

var builtins = []string{"echo", "exit", "type", "pwd", "cd"}

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	reader := bufio.NewReader(os.Stdin)
	var inputBuffer []rune

	for {
		// Clear the current line due to Backspace handling
		fmt.Print("\r$ " + strings.Repeat(" ", len(inputBuffer)+1))
		// Redraw the prompt and input buffer
		fmt.Print("\r$ " + string(inputBuffer))

		r, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}

		switch r {
		case '\t': // TAB pressed
			currentWord := string(inputBuffer)
			var matches []string
			for _, cmd := range builtins {
				if strings.HasPrefix(cmd, currentWord) {
					matches = append(matches, cmd)
				}
			}

			if len(matches) == 0 {
				fmt.Print("\a") // Play bell sound
			} else if len(matches) == 1 {
				inputBuffer = []rune(matches[0] + " ")
			} else if len(matches) > 1 {
				fmt.Println("\n%s", strings.Join(matches, " "))
				inputBuffer = []rune{}
			}

		case '\r', '\n': // Enter pressed
			term.Restore(int(os.Stdin.Fd()), oldState)
			fmt.Println()
			// Restore terminal state before executing the command
			term.Restore(int(os.Stdin.Fd()), oldState)
			handleInput(string(inputBuffer))
			inputBuffer = []rune{}
			// Set terminal back to raw mode after command execution
			oldState, err = term.MakeRaw(int(os.Stdin.Fd()))
			if err != nil {
				panic(err)
			}

		case 127: // Backspace
			if len(inputBuffer) > 0 {
				inputBuffer = inputBuffer[:len(inputBuffer)-1]
			}

		default:
			if r >= 32 && r <= 126 {
				inputBuffer = append(inputBuffer, r)
			}
		}
	}
}

func handleInput(input string) {
	fullCommand := strings.Fields(input)
	if len(fullCommand) == 0 {
		return
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

func run(command string, args []string) {
	var output bytes.Buffer
	cmd := exec.Command(command, args...)
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		fmt.Println(strings.TrimSpace(command) + ": command not found")
		return
	}
	fmt.Print(output.String()) // Print the captured output
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

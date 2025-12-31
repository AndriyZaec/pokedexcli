package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func supportedCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		text := cleanInput(scanner.Text())
		if len(text) <= 0 {
			continue
		}

		cmd, ok := supportedCommands()[strings.ToLower(text[0])]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

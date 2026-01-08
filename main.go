package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AndriyZaec/pokedexcli/internal/api"
	color "github.com/fatih/color"
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
		"map": {
			name:        "map",
			description: "List location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous location areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catching named pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all Pokemons you caught",
			callback:    commandPokedex,
		},
	}
}

func main() {
	printHeader()
	errColor := color.New(color.FgRed)

	client, err := api.NewClient("https://pokeapi.co")
	if err != nil {
		errColor.Println("Network connection error")
	}
	cfg := &config{
		Client: client,
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		color.New(color.FgYellow).Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		text := cleanInput(scanner.Text())
		if len(text) <= 0 {
			continue
		}

		cmd, ok := supportedCommands()[strings.ToLower(text[0])]
		if !ok {
			errColor.Println("Unknown command")
			continue
		}

		cmdErr := cmd.callback(cfg, text[1:]...)
		if cmdErr != nil {
			errColor.Println(cmdErr)
		}
	}
}

func printHeader() {
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Print(
		yellow(`
██████╗  ██████╗ ██╗  ██╗███████╗██████╗ ███████╗██╗  ██╗
██╔══██╗██╔═══██╗██║ ██╔╝██╔════╝██╔══██╗██╔════╝╚██╗██╔╝
██████╔╝██║   ██║█████╔╝ █████╗  ██║  ██║█████╗   ╚███╔╝ 
██╔═══╝ ██║   ██║██╔═██╗ ██╔══╝  ██║  ██║██╔══╝   ██╔██╗ 
██║     ╚██████╔╝██║  ██╗███████╗██████╔╝███████╗██╔╝ ██╗
╚═╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═════╝ ╚══════╝╚═╝  ╚═╝

Welcome to Pokedex CLI
Type `),
		cyan("help"),
		yellow(" to see available commands\n"),
	)
}

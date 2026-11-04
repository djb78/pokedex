package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}


func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
}


func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command := cleanInput(input)
		
		commands := getCommands()
		validCommand, ok := commands[command[0]]	
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := validCommand.callback()
		fmt.Println(err)
	}
}

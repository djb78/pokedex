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
	callback func(*config) error
}

type config struct {
	commands map[string]cliCommand 
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}


func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command := cleanInput(input)
		
		validCommand, ok := cfg.commands[command[0]]	
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := validCommand.callback(cfg)
		fmt.Println(err)
	}
}

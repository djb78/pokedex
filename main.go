package main

func main() {
	endpoint := "https://pokeapi.co/api/v2/location-area/"
	cfg := config{
		commands: map[string]cliCommand {
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
			"map": {
				name: "map",
				description: "",
				callback: commandMap,
			},
			"mapb": {
				name: "mapb",
				description: "",
				callback: commandMapb,
			},
		},
		next: &endpoint,
		previous: nil,
	}

	startRepl(&cfg)
}

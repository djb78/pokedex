package main

import (
	"fmt"
	"github.com/djb78/pokedex/internal/pokeapi"
)

func commandMap(cfg *config) error {
	return commandMapPage(cfg, cfg.next)
}

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		fmt.Println("you are on the first page")
		return nil
	}
	return commandMapPage(cfg, cfg.previous)
}

func commandMapPage(cfg *config, endpoint *string) error {
	client := pokeapi.PokeClient{}
	page, err := client.GetPage(*endpoint)
	if err != nil {
		return err
	}

	cfg.next = page.Next
	cfg.previous = page.Previous
	for _, locArea := range page.Results {
		fmt.Println(locArea.Name)
	}
	return nil
}

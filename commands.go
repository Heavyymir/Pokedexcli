package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Program",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays the next 20 in-game location names",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 in-game location names",
			callback:    commandMapb,
		},
		"explore": {
			name:		 "explore",
			description: "displays the possible encounters in an area",
			callback: 	 commandExplore,
		},
		"catch": {
			name: 		 "catch",
			description: "Attempts to capture a Pokemon",
			callback:	 commandCatch,
		},
		"inspect": {
			name:		 "inspect",
			description: "inspects the data for a Pokemon in the Pokedex",
			callback:	 commandInspect,
		},
		"pokedex": {
			name:		 "pokedex",
			description: "prints a list of all caught pokemon",
			callback:	 commandPokedex,
		},
	}
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println(`Welcome to the Pokedexcli!
Usage:

help: Displays a help message.
exit: Exit the Pokedex.
map: Shows the next 20 in-game locations.
mapb: Shows the previous 20 in-game locations.
explore: Shows the possible encounters in an area.
catch: Attempts to catch a Pokemon.
inspect: Inspects data and statistics for a single caught Pokemon.
pokedex: prints a list of all captured pokemon.`)
	return nil
}

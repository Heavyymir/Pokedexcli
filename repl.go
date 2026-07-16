package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Heavyymir/Pokedex.cli/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	fields := strings.Fields(lowerCase)
	return fields
}

func startRepl() {
	cfg := config{
		Next:          "https://pokeapi.co/api/v2/location-area/",
		Previous:      nil,
		pokeapiClient: pokeapi.NewClient(),
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			return
		}
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			fmt.Println("Empty Command. Please use a valid command. Type 'help' for more information.")
			continue
		}
		command, exists := getCommands()[words[0]]
		if exists {
			args := words[1:]
			err := command.callback(&cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
	return
}

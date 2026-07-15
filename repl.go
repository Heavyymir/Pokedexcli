package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	)

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	fields := strings.Fields(lowerCase)
	return fields 
}


func startRepl() {
	 cfg := config{
              Next:   "https://pokeapi.co/api/v2/location-area/",
              Previous: nil,
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
			err := command.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
	return
}

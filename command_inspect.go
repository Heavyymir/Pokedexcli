package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a Pokemon name")
	}

	name := args[0]

	pokemon, ok := cfg.CaughtPokemon[name]
	if !ok {
		return fmt.Errorf("you have not captured that Pokemon")
	} 
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	for _, element := range pokemon.Types {
		fmt.Println("Types:")
		fmt.Printf("   - %s\n", element.Type.Name)
	}

	return nil
}

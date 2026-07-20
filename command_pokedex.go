package main

import(
	"fmt"
)

func commandPokedex(cfg	*config, args ...string) error {
	for _, entry := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", entry.Name)
	}
	return nil
}

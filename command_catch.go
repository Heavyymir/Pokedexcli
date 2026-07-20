package main

import(
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a Pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	const baseCatch = 20
	chance := float64(baseCatch) / float64(pokemon.BaseExperience)

	catchVal := rand.Float64()
	if catchVal <= chance {
		fmt.Printf("%s successfully caught!\n", pokemon.Name)
		fmt.Printf("Adding %s's data to the Pokedex!\n", pokemon.Name)	
		cfg.CaughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s broke free!\n", pokemon.Name)
	}

	return nil
}

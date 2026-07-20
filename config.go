package main

import "github.com/Heavyymir/Pokedex.cli/internal/pokeapi"

type config struct {
	Next          string
	Previous      *string
	pokeapiClient *pokeapi.Client
	CaughtPokemon map[string]pokeapi.Pokemon	
}

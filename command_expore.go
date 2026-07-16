package main

import (
	"fmt"
	
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Expecting location arguments")
	}
	locationName := args[0]
	areas, err := cfg.pokeapiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}

	for _, encounter := range areas.Encounters {
		fmt.Println(encounter)
	}
	
	return nil
}

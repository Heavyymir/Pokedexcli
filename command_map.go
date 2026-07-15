package main

import "fmt"

func commandMap(cfg *config) error {
	areas, err := getLocationAreas(cfg.Next)
	if err != nil {
		return err
	}

	//Print area results from unmarshaled json data
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = areas.Next
	// Store pagination URLs for future map/mapb calls
	cfg.Previous = areas.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	areas, err := getLocationAreas(*cfg.Previous)
	if err != nil {
		return err
	}

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = areas.Next
	cfg.Previous = areas.Previous

	return nil
}

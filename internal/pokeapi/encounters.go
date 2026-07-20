package pokeapi


type PokemonEncounter struct {
	Pokemon struct {
		Name		string `json:"name"`
		URL			string `json:"url"`
	} `json:"pokemon"`
}

type Pokemon struct {
	Name			string	`json:"name"`
	BaseExperience	int		`json:"base_experience"`
}

package pokeapi


type Pokemon struct {
	Pokemon struct {
		Name		string `json:"name"`
		URL			string `json:"url"`
	} `json:"pokemon"`
}

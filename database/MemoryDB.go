package database

type Pokemon struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
	Type string `json:"Type"`
}

var PokemonDb = []Pokemon{
	Pokemon{ID: "1", Name: "Pikachu", Type: "Electric"},
	Pokemon{ID: "2", Name: "Charmeleon", Type: "Fire"},
	Pokemon{ID: "3", Name: "Arcanine", Type: "Fire"},
}

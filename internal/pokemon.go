package beerscli

// Pokemon representation of pokemon into data struct
type Pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

// PokemonRepo definition of method to access a data pokemon
type PokemonRepo interface {
	GetPokemon(id string) (Pokemon, error)
}

// NewPokemon initialize struct pokemon
func NewPokemon(id int, name string, height, weight int) (p Pokemon) {
	p = Pokemon{
		ID:     id,
		Name:   name,
		Height: height,
		Weight: weight,
	}
	return
}

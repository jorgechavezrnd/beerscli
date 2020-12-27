package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beerscli "github.com/jorgechavezrnd/test_project/internal"
)

const (
	pokemonEndpoint = "pokemon/"
	pokeapiURL      = "https://pokeapi.co/api/v2/"
)

type pokemonRepo struct {
	url string
}

// NewPokeapiRepository fetch beers data from csv
func NewPokeapiRepository() beerscli.PokemonRepo {
	return &pokemonRepo{url: pokeapiURL}
}

func (p *pokemonRepo) GetPokemon(id string) (pokemon beerscli.Pokemon, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v%v", p.url, pokemonEndpoint, id))
	if err != nil {
		return pokemon, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error on read all: ", err)
		return pokemon, err
	}

	err = json.Unmarshal(contents, &pokemon)
	if err != nil {
		fmt.Println("Error on unmarshal: ", err)
		return pokemon, err
	}
	return
}

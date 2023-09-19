package out

import (
	"io"
	"net/http"
)

type PokedexAdapter struct {
}

func NewPokedexAdapter() PokedexAdapterI {
	return &PokedexAdapter{}
}

func (p PokedexAdapter) GetPokemonByName(name string) (interface{}, error) {
	resp, err := http.Get("https://pokeapi.glitch.me/v1/pokemon/snorlax")

	if err != nil {
		return nil, err
		// handle error
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return string(body), err
}

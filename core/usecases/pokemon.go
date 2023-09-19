package usecases

import (
	"pokemon-golang/adapter/out"
	"pokemon-golang/core/domain"
	"pokemon-golang/core/ports"
)

type PokemonUseCase struct {
}

func NewPokemonUseCase() ports.PokemonUseCase {
	return &PokemonUseCase{}
}

func (p PokemonUseCase) SavePokemon(pokemons []*domain.PokemonRequest, pokedexAdapter *out.PokedexAdapter) (interface{}, error) {
	resp, err := pokedexAdapter.GetPokemonByName("pikachu")

	if err != nil {
		println(resp)
	}

	return resp, nil
}

func (p PokemonUseCase) GetAllPokemons() (interface{}, error) {

	return nil, nil
}

package usecases

import (
	"pokemon-golang/core/domain"
	"pokemon-golang/core/ports"
)

type PokemonUseCase struct {
}

func NewPokemonUseCase() ports.PokemonUseCase {
	return &PokemonUseCase{}
}

func (p PokemonUseCase) SavePokemon(pokemons []*domain.PokemonRequest) (interface{}, error) {
	return pokemons, nil
}

func (p PokemonUseCase) GetAllPokemons() (interface{}, error) {

	return nil, nil
}

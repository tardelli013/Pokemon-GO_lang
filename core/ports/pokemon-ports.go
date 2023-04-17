package ports

import "pokemon-golang/core/domain"

type PokemonUseCase interface {
	SavePokemon(pokemons []*domain.PokemonRequest) (interface{}, error)
	GetAllPokemons() (interface{}, error)
}

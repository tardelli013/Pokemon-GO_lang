package ports

import (
	"pokemon-golang/adapter/out"
	"pokemon-golang/core/domain"
)

type PokemonUseCase interface {
	SavePokemon(pokemons []*domain.PokemonRequest, pokedexAdapter out.PokedexAdapter) (interface{}, error)
	GetAllPokemons() (interface{}, error)
}

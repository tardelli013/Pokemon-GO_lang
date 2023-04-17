package adapter

import "pokemon-golang/core/domain"

type Scan interface {
	ScanOperations() ([]*domain.PokemonRequest, error)
}

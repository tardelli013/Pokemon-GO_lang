package out

type PokedexAdapterI interface {
	GetPokemonByName(name string) (interface{}, error)
}

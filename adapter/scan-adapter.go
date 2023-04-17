package adapter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"pokemon-golang/core/domain"
)

type ScanAdapter struct {
}

func NewScan() Scan {
	return &ScanAdapter{}
}

func (s ScanAdapter) ScanOperations() ([]*domain.PokemonRequest, error) {
	var operations []*domain.PokemonRequest

	fmt.Println("\n[SCAN] Enter a operations json: ")
	reader := bufio.NewReader(os.Stdin)
	return decode(reader, operations)
}

func decode(reader *bufio.Reader, oper []*domain.PokemonRequest) ([]*domain.PokemonRequest, error) {
	err := json.NewDecoder(reader).Decode(&oper)
	if err != nil {
		return nil, err
	}
	return oper, nil
}

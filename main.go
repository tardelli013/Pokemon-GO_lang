package main

import (
	"fmt"
	"pokemon-golang/adapter"
	"pokemon-golang/core/ports"
	"pokemon-golang/core/usecases"
	_ "pokemon-golang/docs"
	"pokemon-golang/helpers"
	"time"
)

// @title           Pokemon GO lang
// @version         1.0
// @description     Example of ports and adapters architecture with Golang.
// @termsOfService  https://google.com
// @contact.name   Tardelli Moura
// @contact.url    https://google.com
// @contact.email  tardelli.m@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	useCasePokemon := usecases.NewPokemonUseCase()

	// up gin server in a new Goroutine
	go runWithGinHttp(useCasePokemon)

	// this delay is for gin server up
	time.Sleep(2 * time.Second)
	runWithTerminalScan(useCasePokemon)
}

func runWithGinHttp(useCasePokemon ports.PokemonUseCase) {
	g := adapter.SetupRouter(useCasePokemon)
	err := g.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func runWithTerminalScan(useCase ports.PokemonUseCase) {
	scan := adapter.NewScan()

	operationsFromScan, err := scan.ScanOperations()
	if err != nil {
		panic(err)
	}

	//mudar depois
	print(operationsFromScan)

	feeResults, err := useCase.GetAllPokemons()
	if err != nil {
		panic(err)
	}

	fmt.Println("\n[SCAN] Result: ")
	helpers.PrettyPrint(&feeResults)
	fmt.Print("\n")

	runWithTerminalScan(useCase)
}

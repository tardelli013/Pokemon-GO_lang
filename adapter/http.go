package adapter

import (
	"net/http"
	"pokemon-golang/core/domain"
	"pokemon-golang/core/ports"
	_ "pokemon-golang/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(pokemonUseCase ports.PokemonUseCase) *gin.Engine {
	g := gin.Default()
	v1 := g.Group("/api/v1")

	healthz(v1)
	findPokemonByNameAndSave(pokemonUseCase, v1)
	getAllPokemons(pokemonUseCase, v1)
	swagger(g)

	return g
}

func swagger(g *gin.Engine) gin.IRoutes {
	return g.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// @Summary      healthz
// @Schemes
// @Description  Responds running
// @Tags         healthZ
// @Produce      json
// @Success      200
// @Router       /healthz [get]
func healthz(g *gin.RouterGroup) gin.IRoutes {
	return g.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "running")
	})
}

// @Summary      Find pokemon by name and save in database
// @Schemes
// @Description  Find pokemon by name and save in database
// @Tags         Pokemon
// @Produce      json
// @Param        operations  body      []domain.PokemonRequest  true  "Operations JSON"
// @Success      200   {array}  domain.PokemonRequest
// @Router       /pokemon [post]
func findPokemonByNameAndSave(useCase ports.PokemonUseCase, g *gin.RouterGroup) gin.IRoutes {
	return g.POST("/pokemon", func(c *gin.Context) {
		var input []*domain.PokemonRequest

		err := c.ShouldBindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := useCase.SavePokemon(input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, result)
	})
}

// @Summary      Get all pokemons in database
// @Schemes
// @Description  Get all pokemons in database
// @Tags         Pokemon
// @Produce      json
// @Success      200   {array}  domain.PokemonRequest
// @Router       /pokemon [get]
func getAllPokemons(useCase ports.PokemonUseCase, g *gin.RouterGroup) gin.IRoutes {
	return g.GET("/pokemon", func(c *gin.Context) {

		result, err := useCase.GetAllPokemons()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, result)
	})
}

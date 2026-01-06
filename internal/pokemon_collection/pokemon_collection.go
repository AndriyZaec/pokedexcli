package pokemoncollection

import (
	"math/rand"

	"github.com/AndriyZaec/pokedexcli/internal/api"
)

var collection = map[string]api.Pokemon{}

func AddPokemon(pokemon api.Pokemon) {
	collection[pokemon.Name] = pokemon
}

func GetPokemon(name string) (api.Pokemon, bool) {
	pokemon, ok := collection[name]
	return pokemon, ok
}

func CatchPokemon(pokemon *api.Pokemon) (*api.Pokemon, bool) {
	isCatched := catchChance(pokemon.BaseExperience)

	return pokemon, isCatched
}

func catchChance(baseExp int) bool {
	const (
		A         = 103.125 // hyperbola scale
		B         = 87.5    // hyperbola shift
		minChance = 0.10    // 10%
		maxChance = 0.85    // 85%
	)

	p := A / (float64(baseExp) + B)

	// clamp
	if p < minChance {
		p = minChance
	}
	if p > maxChance {
		p = maxChance
	}

	roll := rand.Float64() // 0.0 .. 1.0
	return roll < p
}

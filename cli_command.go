package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/AndriyZaec/pokedexcli/internal/api"
	pokemoncollection "github.com/AndriyZaec/pokedexcli/internal/pokemon_collection"
)

type config struct {
	Client   *api.Client
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func commandExit(cfg *config, args ...string) error {
	_ = cfg
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	_ = cfg
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for k, v := range supportedCommands() {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func commandMap(cfg *config, args ...string) error {
	urlToFetch := api.FirstLocationAreasURL
	if cfg.Next != nil {
		urlToFetch = *cfg.Next
	}

	resp, err := cfg.Client.GetLocationAreasPageByURL(urlToFetch)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	return nil
}

func commandMapB(cfg *config, args ...string) error {
	urlToFetch := api.FirstLocationAreasURL
	if cfg.Previous != nil {
		urlToFetch = *cfg.Previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := cfg.Client.GetLocationAreasPageByURL(urlToFetch)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no location selected")
	}

	resp, err := cfg.Client.GetLocationInfo(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", resp.Location.Name)
	if len(resp.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
		for _, encounter := range resp.PokemonEncounters {
			fmt.Println(" - ", encounter.Pokemon.Name)
		}
	} else {
		fmt.Println("No Pokemon found :(")
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no pokemon to catch")
	}
	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	resp, err := cfg.Client.GetPokemon(name)
	if err != nil {
		return err
	}

	_, isCatched := pokemoncollection.CatchPokemon(resp)

	if isCatched {
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}

func commandInspect(cfg *config, args ...string) error {
	_ = cfg
	if len(args) < 1 {
		return errors.New("no pokemon name to inspect")
	}

	name := args[0]
	inspectedPokemon, ok := pokemoncollection.InspectPokemon(name)
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println(inspectedPokemon.FormatPokemonInfo())

	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	_ = cfg
	if len(args) > 0 {
		return errors.New("invalid arguments")
	}

	pokemons := pokemoncollection.Pokedex()
	if len(pokemons) < 1 {
		return errors.New("your pokedex is empty... for now")
	}

	fmt.Println("Your Pokedex:")
	for _, p := range pokemons {
		fmt.Println(" - ", p.Name)
	}

	return nil
}

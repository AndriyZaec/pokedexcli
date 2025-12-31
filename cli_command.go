package main

import (
	"fmt"
	"os"

	"github.com/AndriyZaec/pokedexcli/internal/api"
)

type config struct {
	Client   *api.Client
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func commandExit(cfg *config) error {
	_ = cfg
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	_ = cfg
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for k, v := range supportedCommands() {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func commandMap(cfg *config) error {
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

func commandMapB(cfg *config) error {
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

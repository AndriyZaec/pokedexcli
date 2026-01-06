package api

import "fmt"

const (
	LocationAreaURL       = "https://pokeapi.co/api/v2/location-area"
	FirstLocationAreasURL = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	PokemonURL            = "https://pokeapi.co/api/v2/pokemon"
)

func (c *Client) GetLocationAreasPageByURL(fullURL string) (*PaginatedResponse[LocationArea], error) {
	var out PaginatedResponse[LocationArea]
	if err := c.getByURL(fullURL, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) GetLocationInfo(location string) (*LocationInfo, error) {
	var out LocationInfo
	fullURL := fmt.Sprintf("%s/%s", LocationAreaURL, location)
	if err := c.getByURL(fullURL, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) GetPokemon(name string) (*Pokemon, error) {
	var out Pokemon
	fullURL := fmt.Sprintf("%s/%s", PokemonURL, name)
	if err := c.getByURL(fullURL, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

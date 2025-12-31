package api

const FirstLocationAreasURL = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

func (c *Client) GetLocationAreasPageByURL(fullURL string) (*PaginatedResponse[LocationArea], error) {
	var out PaginatedResponse[LocationArea]
	if err := c.getByURL(fullURL, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

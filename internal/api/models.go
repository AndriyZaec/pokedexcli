package api

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PaginatedResponse[T any] struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []T     `json:"results"`
}

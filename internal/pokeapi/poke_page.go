package pokeapi

import (
	"encoding/json"
	"io"
)

type PokePage struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
} 

func decodePokePage(body io.Reader) (PokePage, error) {
	var page PokePage
	err := json.NewDecoder(body).Decode(&page)
	if err != nil {
		return page, err
	}
	return page, nil
}

package pokeapi

import (
	"net/http"
)

type PokeClient struct {

}

func (c PokeClient) GetPage(endpoint string) (PokePage, error) {
	var page PokePage
	res, err := http.Get(endpoint)
	if err != nil {
		return page, err
	}
	defer res.Body.Close()
	
	page, err = decodePokePage(res.Body)
	if err != nil {
		return page, err
	}

	return page, nil
}



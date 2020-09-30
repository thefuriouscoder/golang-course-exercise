package punk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/thefuriouscoder/golang-exercise/internal/model"
)

const (
	punkURL          = "https://api.punkapi.com/v2"
	productsEndpoint = "/beers"
	searchEndpoint   = "/beers"
)

type punkRepo struct {
	url string
}

// NewPunkRepository fetch beers data from Punk Brewery API
func NewPunkRepository() model.PunkRepo {
	return &punkRepo{url: punkURL}
}

func (b *punkRepo) GetBeers() (beers []model.Beer, err error) {
	var url = fmt.Sprintf("%v%v", b.url, productsEndpoint)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, err
	}
	return
}

func (b *punkRepo) GetBeer(id int) (beers []model.Beer, err error) {
	var url = fmt.Sprintf("%v%v/%v", b.url, productsEndpoint, id)
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, err
	}

	return
}

func (b *punkRepo) Search(terms map[string]string) (beers []model.Beer, err error) {
	var baseURL = fmt.Sprintf("%v%v", b.url, productsEndpoint)
	var searchURL = getSearchURL(baseURL, terms)

	fmt.Println(searchURL)
	response, err := http.Get(searchURL)
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, err
	}
	return

}

func getSearchURL(baseURL string, terms map[string]string) string {

	params := url.Values{}
	for k, v := range terms {
		params.Add(k, v)
	}

	return fmt.Sprintf("%v?%v", baseURL, params.Encode())

}

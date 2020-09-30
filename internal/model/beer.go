package model

// Beer representation of beer into data struct
type Beer struct {
	ProductID   int     `json:"id"`
	Name        string  `json:"name"`
	Tagline     string  `json:"tagline"`
	Description string  `json:"description"`
	ABV         float64 `json:"abv"`
	IBU         float64 `json:"ibu"`
}

// Terms struct for searching beers by name, taglin, description, abv or ibu
type Terms struct {
	Name   string  `url:"beer_name"`
	MinAbv float64 `url:"abv_gt"`
	MaxAbv float64 `url:"abv_lt"`
	MinIbu float64 `url:"ibu_gt"`
	MaxIbu float64 `url:"ibu_lt"`
	Yeast  string  `url:"yeast"`
	Malt   string  `url:"malt"`
	Hops   string  `url:"hops"`
}

// PunkRepo definiton of methods to access a data beer
type PunkRepo interface {
	GetBeers() ([]Beer, error)
	GetBeer(id int) ([]Beer, error)
	Search(map[string]string) ([]Beer, error)
}

// NewBeer initialize struct beer
func NewBeer(id int, name, tagline, description string, abv, ibu float64) (b Beer) {
	b = Beer{
		ProductID:   id,
		Name:        name,
		Tagline:     tagline,
		Description: description,
		ABV:         abv,
		IBU:         ibu,
	}
	return
}

// NewTerms initializes struct terms
func NewTerms(name, yeast, malt, hops string, minIbu, maxIbu, minAbv, maxAbv float64) (t Terms) {
	t = Terms{
		Name:   name,
		MinAbv: minAbv,
		MaxAbv: maxAbv,
		MinIbu: minIbu,
		MaxIbu: maxIbu,
		Yeast:  yeast,
		Hops:   hops,
		Malt:   malt,
	}
	return
}

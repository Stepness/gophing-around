package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	Id        int `json:"id"`
	Name      string
	CreatedOn string
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p) //Using encoder over writing the response avoids allocating memory to generate the json
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	{
		Id:        1,
		Name:      "Latte",
		CreatedOn: time.Now().UTC().String(),
	},
	{
		Id:        2,
		Name:      "Coffee",
		CreatedOn: time.Now().UTC().String(),
	},
}

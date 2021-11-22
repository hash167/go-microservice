package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Price       float32 `json:"price" validate:"gt=0"`
	Description string  `json:"description"`
	SKU         string  `json:"sku" validate:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

// Method on type Product and not Products
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	e := d.Decode(i)
	return e
}

func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

func GetProducts() Products {
	return productList
}

func GetProductById(id int) (Product, error) {
	p, _, err := findProduct(id)
	return *p, err

}

func AddProduct(p *Product) {
	p.ID = getProductID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, i, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[i] = p
	return nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	_, i, _ := findProduct(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

// Explicit error
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getProductID() int {
	return len(productList) + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

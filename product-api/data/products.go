package data

import (
	"fmt"
	"time"
)

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

func GetProductById(id int) (*Product, error) {
	i := findIndexByProductId(id)
	if i == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

func UpdateProduct(p Product) error {
	i := findIndexByProductId(p.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	productList[i] = &p
	return nil
}

func AddProduct(p Product) {
	maxId := productList[len(productList)-1].ID
	p.ID = maxId + 1
	productList = append(productList, &p)
}

func DeleteProduct(id int) error {
	i := findIndexByProductId(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1:]...)
	return nil
}

func findIndexByProductId(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

// productList is a hard coded list of products for this
// example data source
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

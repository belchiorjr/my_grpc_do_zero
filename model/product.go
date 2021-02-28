package model

import "github.com/satori/go.uuid"

// Product estrutura de produtos
type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewProduct método responsavel por retornar uma instancia de produto com id
func NewProduct() *Product {
	return &Product{
		ID: uuid.NewV4().String(),
	}
}

// Products listagem de produtos 
type Products struct {
	Product []*Product `json:"products"`
}

// Add pertence a estrutura de products, adiciona uma struct de produtos
func (p *Products) Add(product *Product) {
	p.Product = append(p.Product, product)
}

// NewProducts retorna a listagem de produtos adicionados
func NewProducts() *Products {
	return &Products{}
}

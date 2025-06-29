package products

import "query-service/internal/domain/models/categories"

type Product struct {
	id       string
	name     string
	price    uint32
	category categories.Category
}

func NewProduct(id, name string, price uint32, category categories.Category) *Product {
	return &Product{
		id:       id,
		name:     name,
		price:    price,
		category: category,
	}
}

func (p *Product) Id() string {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Price() uint32 {
	return p.price
}

func (p *Product) Category() categories.Category {
	return p.category
}

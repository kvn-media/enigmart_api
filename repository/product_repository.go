package repository

import (
	"git.enigmacamp.com/enigmart-api/models"
)

type ProductRepository interface {
	Insert(newProduct *models.Product) models.Product
	List() []models.Product
	// Id(id string) (models.Product, error)
}

type productRepository struct {
	db []models.Product
}

func (p *productRepository) Insert(newProduct *models.Product) models.Product {
	p.db = append(p.db, *newProduct)
	return *newProduct
}

func (p *productRepository) List() []models.Product {
	return p.db
}

// func (p *productRepository) Id(id string) (models.Product, error)  {
// 	var product models.Product
// 	if err := p.db.Get(&product, id); err != nil {
// 		return models.Product{}, err
// 	}
// 	return product, nil
// }

func NewProductRepository() ProductRepository {
	repo := new(productRepository)
	return repo
}

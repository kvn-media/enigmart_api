package repository

import (
	"git.enigmacamp.com/enigmart-api/models"
)

type ProductRepository interface {
	Insert(newProduct *models.Product) models.Product
	List() []models.Product
	Update(newProduct *models.Product) models.Product
	Get(id string) models.Product
	Delete(id string) bool
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

func (p *productRepository) Update(newProduct *models.Product) models.Product  {
	var productRes models.Product
	for i := 0; i < len(p.db); i++ {
		if p.db[i].Id == newProduct.Id {
			p.db[i] = *newProduct
			productRes = p.db[i]
		}
	}
	return productRes
}

func (p *productRepository) Get(id string) models.Product {
	var product models.Product
	for _, item := range p.db {
		if item.Id == id {
			product = item
		}
	}
	return product
}

func (p *productRepository) Delete(id string) bool {
	var products []models.Product
	result := false
	for i := 0; i < len(p.db); i++ {
		if p.db[i].Id == id {
			products = append(p.db[:i], p.db[i+1:]...)
			result = true
		}
	}
	p.db = products
	return result
}

func NewProductRepository() ProductRepository {
	repo := new(productRepository)
	return repo
}

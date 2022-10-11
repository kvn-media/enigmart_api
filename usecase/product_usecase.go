package usecase

import (
	"git.enigmacamp.com/enigmart-api/models"
	"git.enigmacamp.com/enigmart-api/repository"
)

type ProductUseCase interface {
	CreateNewProduct(newProduct *models.Product) models.Product
	GetAllProduct() []models.Product
	// GetProduct() []models.Product
}

type productUseCase struct {
	repo repository.ProductRepository
}

func (p *productUseCase) CreateNewProduct(newProduct *models.Product) models.Product {
	return p.repo.Insert(newProduct)
}

func (p *productUseCase) GetAllProduct() []models.Product {
	return p.repo.List()
}

// func (p *productUseCase) GetProduct() []models.Product {
// 	return p.repo.Id()
// }

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	pc := new(productUseCase)
	pc.repo = repo
	return pc
}

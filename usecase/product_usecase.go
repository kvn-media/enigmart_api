package usecase

import (
	"git.enigmacamp.com/enigmart-api/models"
	"git.enigmacamp.com/enigmart-api/repository"
)

type ProductUseCase interface {
	CreateNewProduct(newProduct *models.Product) models.Product
	GetAllProduct() []models.Product
	UpdateProduct(newProduct models.Product) models.Product
	DeleteProduct(id string)bool
	GetProductById(id string)models.Product
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

func (p *productUseCase) UpdateProduct(newProduct models.Product) models.Product{
	return p.repo.Update(&newProduct)
}

func (p *productUseCase) DeleteProduct(id string)bool {
	return p.repo.Delete(id)
}

func (p *productUseCase) GetProductById(id string)models.Product{
	return p.repo.Get(id)
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	pc := new(productUseCase)
	pc.repo = repo
	return pc
}

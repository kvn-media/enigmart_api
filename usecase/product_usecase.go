package usecase

import (
	"git.enigmacamp.com/enigmart-api/models"
	"git.enigmacamp.com/enigmart-api/repository"
	"git.enigmacamp.com/enigmart-api/utils"
)

type ProductUseCase interface {
	CreateNewProduct(newProduct *models.Product) error
	GetProductById(id string) (models.Product, error)
	GetAllProduct(page int, totalRows int) ([]models.Product, error)
	UpdateProduct(newProduct models.Product) error
	DeleteProduct(id string) error
}

type productUseCase struct {
	repo repository.ProductRepository
}

func (p *productUseCase) CreateNewProduct(newProduct *models.Product) error {
	newProduct.Id = utils.GenerateId()
	return p.repo.Insert(newProduct)
}

func (p *productUseCase) GetProductById(id string) (models.Product, error) {
	return p.repo.Get(id)
}

func (p *productUseCase) GetAllProduct(page int, totalRows int) ([]models.Product, error) {
	return p.repo.List(page, totalRows)
}

func (p *productUseCase) UpdateProduct(newProduct models.Product) error {
	return p.repo.Update(&newProduct)
}

func (p *productUseCase) DeleteProduct(id string) error {
	return p.repo.Delete(id)
}


func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	pc := new(productUseCase)
	pc.repo = repo
	return pc
}

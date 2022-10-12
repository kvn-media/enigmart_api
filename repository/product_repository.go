package repository

import (
	"git.enigmacamp.com/enigmart-api/models"
	"git.enigmacamp.com/enigmart-api/utils"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Insert(newProduct *models.Product) error
	Get(id string) (models.Product,error)
	List(page int, totalRows int) ([]models.Product, error)
	Update(newProduct *models.Product) error
	Delete(id string) error
}

type productRepository struct {
	db *sqlx.DB
}

func (p *productRepository) Insert(newProduct *models.Product) error {
	_, err := p.db.NamedExec(utils.INSERT_PRODUCT, newProduct)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) Get(id string) (models.Product,error) {
	var product models.Product
	err := p.db.Get(&product,utils.SELECT_PRODUCT_ID, id)
	if err != nil {
		return models.Product{},err
	}
	return product,nil
}

func (p *productRepository) List(page int, totalRows int) ([]models.Product, error) {
	limit := totalRows
	offset := limit * (page - 1)
	var products []models.Product
	err := p.db.Select(&products,utils.SELECT_PRODUCT_LIST, limit, offset)
	if err != nil {
		return nil, err
	}
	return products,nil
}

func (p *productRepository) Update(product *models.Product) error {
	_, err := p.db.NamedExec(utils.UPDATE_PRODUCT, product)
	if err != nil {
		return err
	}
	return nil
}


func (p *productRepository) Delete(id string) error {
	_, err := p.db.Exec(utils.DELETE_PRODUCT, id)
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}

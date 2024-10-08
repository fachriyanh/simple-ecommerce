package repository

import (
	"go-hypefast/entity"

	"gorm.io/gorm"
)

type Usecase interface {
	CreateProduct(product *entity.Product) error
	GetProducts() (map[string]*entity.Product, error)
	CreateOrder(order *entity.Order, product *entity.Product) error
	ListOrder() ([]*entity.Order, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateProduct(product *entity.Product) error {
	entity.ListProduct[product.SKU] = product
	return nil
}

func (r *Repository) GetProducts() (map[string]*entity.Product, error) {
	if entity.ListProduct == nil {
		return nil, entity.ErrorNotFound
	}

	return entity.ListProduct, nil
}

func (r *Repository) CreateOrder(order *entity.Order, product *entity.Product) error {
	product.Stock -= order.Items
	entity.ListOrder = append(entity.ListOrder, &entity.Order{
		Name:  order.Name,
		SKU:   order.SKU,
		Items: order.Items,
	})
	return nil
}

func (r *Repository) ListOrder() ([]*entity.Order, error) {
	return entity.ListOrder, nil
}

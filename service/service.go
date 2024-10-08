package service

import (
	"go-hypefast/entity"
	"go-hypefast/repository"
)

type Usecase interface {
	CreateProduct(product *entity.Product) error
	GetProducts() (map[string]*entity.Product, error)
	CreateOrder(order *entity.Order) error
	ListOrder() ([]*entity.Order, error)
}

type Service struct {
	repository repository.Usecase
}

func NewService(repository repository.Usecase) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateProduct(product *entity.Product) error {
	// Additional validation post live coding : add validation SKU already created
	if product := entity.ListProduct[product.SKU]; product != nil {
		return entity.ErrorConflict
	}

	return s.repository.CreateProduct(product)
}

func (s *Service) GetProducts() (map[string]*entity.Product, error) {
	return s.repository.GetProducts()
}

func (s *Service) CreateOrder(order *entity.Order) error {
	product := entity.ListProduct[order.SKU]

	// Additional validation post live coding : add validation SKU Not Found
	if product == nil {
		return entity.ErrorNotFound
	}

	if entity.ListProduct[order.SKU].Stock < order.Items {
		return entity.ErrorNotAvailable
	}

	return s.repository.CreateOrder(order, entity.ListProduct[order.SKU])
}

func (s *Service) ListOrder() ([]*entity.Order, error) {
	return s.repository.ListOrder()
}

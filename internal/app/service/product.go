package service

import (
	"github.com/gustavodalves/ecommerce/internal/domain/models"
	"github.com/gustavodalves/ecommerce/internal/domain/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func (ps *ProductService) Create(name string, price float64) error {
	product := models.NewProduct(name, price)

	err := ps.ProductRepository.Save(product)

	if err != nil {
		return err
	}

	return nil
}

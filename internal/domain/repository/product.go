package repository

import "github.com/gustavodalves/ecommerce/internal/domain/models"

type ProductRepository interface {
	Save(*models.Product) error
	FindById(id string) *models.Product
}

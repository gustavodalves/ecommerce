package repository

import "github.com/gustavodalves/ecommerce/internal/domain/models"

type OrderRepository interface {
	Save(o *models.Order) error
	FindByID(id string) (*models.Order, error)
}

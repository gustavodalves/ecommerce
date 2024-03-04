package service

import (
	"github.com/gustavodalves/ecommerce/internal/domain/models"
	"github.com/gustavodalves/ecommerce/internal/domain/repository"
)

type OrderService struct {
	Repository repository.OrderRepository
}

func (os *OrderService) CreateOrder() error {
	order := models.NewOrder()
	err := os.Repository.Save(order)

	if err != nil {
		return err
	}

	return nil
}

func (os *OrderService) FindByID(id string) (*models.Order, error) {
	order, err := os.Repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (os *OrderService) Approve(id string) error {
	order, err := os.Repository.FindByID(id)

	if err != nil {
		return err
	}

	err = order.ProcessOrder()

	if err != nil {
		return err
	}

	os.Repository.Save(order)

	return nil
}

func (os *OrderService) Reject(id string) error {
	order, err := os.Repository.FindByID(id)

	if err != nil {
		return err
	}

	err = order.CancelOrder()

	if err != nil {
		return err
	}

	os.Repository.Save(order)

	return nil
}

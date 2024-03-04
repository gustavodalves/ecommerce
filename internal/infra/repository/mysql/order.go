package repository_mysql

import (
	"database/sql"
	"errors"

	"github.com/gustavodalves/ecommerce/internal/domain/models"
)

type OrderRepositoryMySQL struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepositoryMySQL {
	return &OrderRepositoryMySQL{DB: db}
}
func (or *OrderRepositoryMySQL) Save(order *models.Order) error {
	tx, err := or.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.Exec("INSERT INTO orders (id, status) VALUES (?, ?)", order.Id, order.Status.GetStatus())
	if err != nil {
		return err
	}

	for _, product := range order.Cart.Products {
		_, err = tx.Exec("INSERT INTO order_products (order_id, product_id) VALUES (?, ?)", order.Id, product.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (or *OrderRepositoryMySQL) FindByID(id string) (*models.Order, error) {
	var orderID, status string
	err := or.DB.QueryRow("SELECT id, status FROM orders WHERE id = ?", id).Scan(&orderID, &status)
	if err != nil {
		return nil, err
	}

	rows, err := or.DB.Query("SELECT p.id, p.name, p.price FROM order_products op LEFT JOIN products p ON p.id = op.product_id WHERE op.order_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	cart := models.RecoverCart(products)

	order := &models.Order{
		Id:     id,
		Cart:   cart,
		Status: mapStatus(status),
	}

	if order.Status == nil {
		return nil, errors.New("unknown order status")
	}

	return order, nil
}

func mapStatus(status string) models.OrderStatus {
	switch status {
	case "created":
		return &models.CreatedStatus{}
	case "pending":
		return &models.PendingStatus{}
	case "approved":
		return &models.ApprovedStatus{}
	case "rejected":
		return &models.RejectedStatus{}
	default:
		return nil
	}
}

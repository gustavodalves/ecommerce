package models

import "errors"

type OrderStatus interface {
	ProcessOrder(o *Order) error
	CancelOrder(o *Order) error
}

type CreatedStatus struct{}

func (cs *CreatedStatus) ProcessOrder(o *Order) error {
	o.Status = &PendingStatus{}
	return nil
}

func (cs *CreatedStatus) CancelOrder(o *Order) error {
	o.Status = &RejectedStatus{}
	return nil
}

type PendingStatus struct{}

func (ps *PendingStatus) ProcessOrder(o *Order) error {
	o.Status = &ApprovedStatus{}
	return nil
}

func (ps *PendingStatus) CancelOrder(o *Order) error {
	o.Status = &RejectedStatus{}
	return nil
}

type ApprovedStatus struct{}

func (as *ApprovedStatus) ProcessOrder(o *Order) error {
	return errors.New("cannot process an already approved order")
}

func (as *ApprovedStatus) CancelOrder(o *Order) error {
	return errors.New("cannot cancel an already approved order")
}

type RejectedStatus struct{}

func (rs *RejectedStatus) ProcessOrder(o *Order) error {
	return errors.New("cannot process a rejected order")
}

func (rs *RejectedStatus) CancelOrder(o *Order) error {
	return errors.New("cannot cancel a rejected order again")
}

type Order struct {
	Cart   *Cart
	Status OrderStatus
}

func NewOrder() *Order {
	return &Order{
		Cart:   &Cart{},
		Status: &CreatedStatus{},
	}
}

func (o *Order) CalculateTotal() float64 {
	return o.Cart.CalculateTotal()
}

func (o *Order) ProcessOrder() error {
	return o.Status.ProcessOrder(o)
}

func (o *Order) CancelOrder() error {
	return o.Status.CancelOrder(o)
}

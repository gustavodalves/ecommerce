package models

type Cart struct {
	Products []*Product
}

func NewCart() *Cart {
	return &Cart{}
}

func RecoverCart(products []*Product) *Cart {
	return &Cart{
		Products: products,
	}
}

func (c *Cart) AddProduct(product *Product) {
	c.Products = append(c.Products, product)
}

func (c *Cart) CalculateTotal() float64 {
	var total float64 = 0

	for _, product := range c.Products {
		total += product.Price
	}

	return total
}

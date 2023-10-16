package usecase

import "github.com/pauloscardoso/go-intensivo-jul/internal/entity"

type OrderInput struct {
	ID    string
	PRICE float64
	TAX   float64
}

type OrderOutput struct {
	ID         string
	PRICE      float64
	TAX        float64
	FinalPrice float64
}

// SOLID - "DIP - Dependency Inversion Principle"
type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPrice(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPrice {
	return &CalculateFinalPrice{
		OrderRepository: orderRepository,
	}
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.ID, input.PRICE, input.TAX)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutput{
		ID:         order.ID,
		PRICE:      order.Price,
		TAX:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}

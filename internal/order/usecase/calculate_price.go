package usecase

import (
	"github.com/eferroni/gointensivo/internal/order/entity"
	"github.com/eferroni/gointensivo/internal/order/infra/database"
)

type OrderInputDto struct {
	ID string
	Price float64
	Tax float64
}

type OrderOutputDto struct {
	ID string
	Price float64
	Tax float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPriceUseCase(orderRepository database.OrderRepository) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{
		OrderRepository: &orderRepository,
	}
}

func (c *CalculateFinalPriceUseCase) Execute(input OrderInputDto) (*OrderOutputDto, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
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
	return &OrderOutputDto{
		ID: order.ID,
		Price: order.Price,
		Tax: order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
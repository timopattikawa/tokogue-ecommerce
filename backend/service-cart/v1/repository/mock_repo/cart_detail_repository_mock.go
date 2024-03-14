package mock_repo

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"service-cart/internal/domain"
)

type CartDetailRepositoryMock struct {
	Mock mock.Mock
}

func (c *CartDetailRepositoryMock) FindByCartIdAndProductId(cartId uuid.UUID, productId int) *domain.CartDetail {
	arguments := c.Mock.Called(cartId, productId)
	if arguments.Get(0) == nil {
		return nil
	}

	return arguments.Get(0).(*domain.CartDetail)
}

func (c *CartDetailRepositoryMock) InsertCartDetail(detail domain.CartDetail) error {
	arguments := c.Mock.Called(detail)
	if arguments.Get(0).(error) == nil {
		return nil
	}

	return arguments.Get(0).(error)
}

func (c *CartDetailRepositoryMock) UpdateCartDetail(detail domain.CartDetail) error {
	arguments := c.Mock.Called(detail)
	if arguments.Get(0).(error) == nil {
		return nil
	}

	return arguments.Get(0).(error)
}

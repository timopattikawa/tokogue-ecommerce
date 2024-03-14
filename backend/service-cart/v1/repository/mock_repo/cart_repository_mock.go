package mock_repo

import (
	"github.com/stretchr/testify/mock"
	"service-cart/internal/domain"
)

type CartRepositoryMock struct {
	Mock mock.Mock
}

func (c *CartRepositoryMock) CreateNewCart(cart domain.Cart) error {
	arguments := c.Mock.Called(cart)
	if arguments.Get(0).(error) == nil {
		return nil
	}

	return arguments.Get(0).(error)
}

func (c *CartRepositoryMock) FindCurrentCart(userId int) *domain.Cart {
	arguments := c.Mock.Called(userId)
	if arguments.Get(0).(*domain.Cart) == nil {
		return nil
	}

	return arguments.Get(0).(*domain.Cart)
}

func (c *CartRepositoryMock) UpdateCart(cart domain.Cart) error {
	arguments := c.Mock.Called(cart)
	if arguments.Get(0).(error) == nil {
		return nil
	}
	return arguments.Get(0).(error)
}

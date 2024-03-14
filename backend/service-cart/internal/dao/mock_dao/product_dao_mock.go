package mock_dao

import (
	"github.com/stretchr/testify/mock"
	"service-cart/internal/dto"
)

type ProductDaoMock struct {
	Mock mock.Mock
}

func (p *ProductDaoMock) StockAvail(request dto.ProductStockRequest) *dto.ProductStockResponse {
	arguments := p.Mock.Called(request)

	if arguments.Get(0) == nil {
		return nil
	}

	return arguments.Get(0).(*dto.ProductStockResponse)
}

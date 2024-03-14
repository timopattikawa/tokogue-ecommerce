package service

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"service-cart/internal/dao/mock_dao"
	"service-cart/internal/domain"
	"service-cart/internal/dto"
	"service-cart/v1/repository/mock_repo"
	"testing"
	"time"
)

func TestAddToCart_CaseCreateNewCart(t *testing.T) {
	mockCartRepo := mock_repo.CartRepositoryMock{Mock: mock.Mock{}}
	mockCartDetailRepo := mock_repo.CartDetailRepositoryMock{Mock: mock.Mock{}}
	mockDaoProduct := mock_dao.ProductDaoMock{Mock: mock.Mock{}}
	cartService := CartServiceImpl{
		cartRepository:   &mockCartRepo,
		detailRepository: &mockCartDetailRepo,
		productDao:       &mockDaoProduct,
	}

	var requestDummy = dto.AddToCartRequest{
		UserId:    1,
		ProductId: 1,
		Qty:       5,
	}

	var dummyProductStockReq = dto.ProductStockRequest{
		ProductId: 1,
		Qty:       5,
	}

	var dummyProductStockRes = dto.ProductStockResponse{
		Price:          10000,
		IsStockReduced: true,
	}

	var dummyCart = domain.Cart{
		Id:          uuid.New(),
		UserId:      1,
		TotalAmount: 0,
		StatusCart:  domain.InBasket,
		CreateAt:    time.Now(),
	}

	var dummyDetailCart = domain.CartDetail{
		CartId:    dummyCart.Id,
		ProductId: requestDummy.ProductId,
		Qty:       requestDummy.Qty,
	}

	mockDaoProduct.Mock.On("StockAvail", dummyProductStockReq).Return(dummyProductStockRes)
	mockCartRepo.Mock.On("FindCurrentCart", requestDummy.UserId).Return(nil)
	mockCartRepo.Mock.On("CreateNewCart", dummyCart).Return(nil)
	mockCartRepo.Mock.On("FindCurrentCart", requestDummy.UserId).Return(dummyCart)
	mockCartDetailRepo.Mock.On("FindByCartIdAndProductId", dummyCart.Id, requestDummy.ProductId).Return(nil)
	mockCartDetailRepo.Mock.On("InsertCartDetail", dummyDetailCart).Return(nil)

	result, err := cartService.AddToCart(requestDummy)
	assert.Equal(t, result.StatusToCart, "SUCCESS")
	assert.NoError(t, err)
}

func TestAddToCart_StockNotAvail(t *testing.T) {
	mockCartRepo := mock_repo.CartRepositoryMock{Mock: mock.Mock{}}
	mockCartDetailRepo := mock_repo.CartDetailRepositoryMock{Mock: mock.Mock{}}
	mockDaoProduct := mock_dao.ProductDaoMock{Mock: mock.Mock{}}
	cartService := CartServiceImpl{
		cartRepository:   &mockCartRepo,
		detailRepository: &mockCartDetailRepo,
		productDao:       &mockDaoProduct,
	}

	var requestDummy = dto.AddToCartRequest{
		UserId:    1,
		ProductId: 1,
		Qty:       5,
	}

	var dummyProductStockReq = dto.ProductStockRequest{
		ProductId: 1,
		Qty:       5,
	}

	var dummyProductStockRes = dto.ProductStockResponse{
		Price:          10000,
		IsStockReduced: false,
	}

	var dummyCart = domain.Cart{
		Id:          uuid.New(),
		UserId:      1,
		TotalAmount: 0,
		StatusCart:  domain.InBasket,
		CreateAt:    time.Now(),
	}

	var dummyDetailCart = domain.CartDetail{
		CartId:    dummyCart.Id,
		ProductId: requestDummy.ProductId,
		Qty:       requestDummy.Qty,
	}

	mockDaoProduct.Mock.On("StockAvail", dummyProductStockReq).Return(dummyProductStockRes)
	mockCartRepo.Mock.On("FindCurrentCart", requestDummy.UserId).Return(nil)
	mockCartRepo.Mock.On("CreateNewCart", dummyCart).Return(nil)
	mockCartRepo.Mock.On("FindCurrentCart", requestDummy.UserId).Return(dummyCart)
	mockCartDetailRepo.Mock.On("FindByCartIdAndProductId", dummyCart.Id, requestDummy.ProductId).Return(nil)
	mockCartDetailRepo.Mock.On("InsertCartDetail", dummyDetailCart).Return(nil)

	result, err := cartService.AddToCart(requestDummy)
	assert.Equal(t, result.StatusToCart, dto.AddToCartResponse{})
	assert.NotNil(t, err)
}

func TestAddToCart_CaseCartExist(t *testing.T) {
	mockCartRepo := mock_repo.CartRepositoryMock{Mock: mock.Mock{}}
	mockCartDetailRepo := mock_repo.CartDetailRepositoryMock{Mock: mock.Mock{}}
	mockDaoProduct := mock_dao.ProductDaoMock{Mock: mock.Mock{}}
	cartService := CartServiceImpl{
		cartRepository:   &mockCartRepo,
		detailRepository: &mockCartDetailRepo,
		productDao:       &mockDaoProduct,
	}

	var requestDummy = dto.AddToCartRequest{
		UserId:    1,
		ProductId: 1,
		Qty:       5,
	}

	var dummyProductStockReq = dto.ProductStockRequest{
		ProductId: 1,
		Qty:       5,
	}

	var dummyProductStockRes = dto.ProductStockResponse{
		Price:          10000,
		IsStockReduced: true,
	}

	var dummyCart = domain.Cart{
		Id:          uuid.New(),
		UserId:      1,
		TotalAmount: 0,
		StatusCart:  domain.InBasket,
		CreateAt:    time.Now(),
	}

	var dummyDetailCart = domain.CartDetail{
		CartId:    dummyCart.Id,
		ProductId: requestDummy.ProductId,
		Qty:       requestDummy.Qty,
	}

	mockDaoProduct.Mock.On("StockAvail", dummyProductStockReq).Return(dummyProductStockRes)
	mockCartRepo.Mock.On("FindCurrentCart", requestDummy.UserId).Return(dummyCart)
	mockCartDetailRepo.Mock.On("FindByCartIdAndProductId", dummyCart.Id, requestDummy.ProductId).Return(nil)
	mockCartDetailRepo.Mock.On("InsertCartDetail", dummyDetailCart).Return(nil)

	result, err := cartService.AddToCart(requestDummy)
	assert.NoError(t, err)
	assert.Equal(t, result.StatusToCart, "SUCCESS")
}

func TestAddToCart_CaseCartExistAndProductExist(t *testing.T) {
	mockCartRepo := mock_repo.CartRepositoryMock{Mock: mock.Mock{}}
	mockCartDetailRepo := mock_repo.CartDetailRepositoryMock{Mock: mock.Mock{}}
	mockDaoProduct := mock_dao.ProductDaoMock{Mock: mock.Mock{}}
	cartService := CartServiceImpl{
		cartRepository:   &mockCartRepo,
		detailRepository: &mockCartDetailRepo,
		productDao:       &mockDaoProduct,
	}

	var requestDummy = dto.AddToCartRequest{
		UserId:    1,
		ProductId: 1,
		Qty:       5,
	}

	var dummyProductStockReq = dto.ProductStockRequest{
		ProductId: 1,
		Qty:       5,
	}

	var dummyProductStockRes = dto.ProductStockResponse{
		Price:          10000,
		IsStockReduced: true,
	}

	var dummyCart = domain.Cart{
		Id:          uuid.New(),
		UserId:      1,
		TotalAmount: 0,
		StatusCart:  domain.InBasket,
		CreateAt:    time.Now(),
	}

	var dummyDetailCart = domain.CartDetail{
		CartId:    dummyCart.Id,
		ProductId: requestDummy.ProductId,
		Qty:       requestDummy.Qty,
	}

	mockDaoProduct.Mock.On("StockAvail", dummyProductStockReq).Return(dummyProductStockRes)
	mockCartRepo.Mock.On("FindCurrentCart", requestDummy.UserId).Return(dummyCart)
	mockCartDetailRepo.Mock.On("FindByCartIdAndProductId", dummyCart.Id, requestDummy.ProductId).Return(dummyDetailCart)
	mockCartDetailRepo.Mock.On("UpdateCartDetail", dummyDetailCart).Return(nil)

	result, err := cartService.AddToCart(requestDummy)
	assert.NoError(t, err)
	assert.Equal(t, result.StatusToCart, "SUCCESS")
}

package service

import (
	"github.com/google/uuid"
	"log"
	"service-cart/internal/dao"
	"service-cart/internal/domain"
	"service-cart/internal/dto"
	"service-cart/internal/exception"
)

type CartServiceImpl struct {
	cartRepository   domain.CartRepository
	detailRepository domain.CartDetailRepository
	productDao       dao.ProductDao
}

type CartService interface {
	GetMyCart(userId int) (dto.MyCartResponse, error)
	AddToCart(request dto.AddToCartRequest) (dto.AddToCartResponse, error)
}

func (c CartServiceImpl) GetMyCart(userId int) (dto.MyCartResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CartServiceImpl) AddToCart(request dto.AddToCartRequest) (dto.AddToCartResponse, error) {
	requestCheck := dto.ProductStockRequest{
		ProductId: request.ProductId,
		Qty:       request.Qty,
	}
	availResponse := c.productDao.StockAvail(requestCheck)

	if !availResponse.IsStockReduced {
		return dto.AddToCartResponse{}, exception.BadRequestError{Message: "Stock Not Available"}
	}

	cart := c.cartRepository.FindCurrentCart(request.UserId)
	if cart == nil {
		var cart = domain.Cart{
			Id:          uuid.New(),
			UserId:      request.UserId,
			TotalAmount: 0,
		}
		err := c.cartRepository.CreateNewCart(cart)
		if err != nil {
			log.Println("Error Create Cart")
			return dto.AddToCartResponse{}, err
		}
		//cart = c.cartRepository.FindCurrentCart(request.UserId)
	}
	cart.TotalAmount = (request.Qty * availResponse.Price)
	err := c.cartRepository.UpdateCart(*cart)
	if err != nil {
		return dto.AddToCartResponse{}, err
	}
}

func NewCartService(
	cartRepository domain.CartRepository,
	detailRepository domain.CartDetailRepository,
	productDao dao.ProductDao) CartService {
	return &CartServiceImpl{
		cartRepository:   cartRepository,
		detailRepository: detailRepository,
		productDao:       productDao,
	}
}

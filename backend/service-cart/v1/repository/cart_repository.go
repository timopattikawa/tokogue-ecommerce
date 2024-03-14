package repository

import (
	"gorm.io/gorm"
	"service-cart/internal/domain"
)

type CartRepositoryImpl struct {
	db *gorm.DB
}

func (c CartRepositoryImpl) CreateNewCart(cart domain.Cart) error {
	//TODO implement me
	panic("implement me")
}

func (c CartRepositoryImpl) FindCurrentCart(userId int) *domain.Cart {
	//TODO implement me
	panic("implement me")
}

func (c CartRepositoryImpl) UpdateCart(cart domain.Cart) error {
	//TODO implement me
	panic("implement me")
}

func NewCartRepository(db *gorm.DB) domain.CartRepository {
	return &CartRepositoryImpl{
		db: db,
	}
}

package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"service-cart/internal/domain"
)

type CartDetailRepositoryImpl struct {
	db *gorm.DB
}

func (c CartDetailRepositoryImpl) InsertCartDetail(detail domain.CartDetail) error {
	//TODO implement me
	panic("implement me")
}

func (c CartDetailRepositoryImpl) UpdateCartDetail(detail domain.CartDetail) error {
	//TODO implement me
	panic("implement me")
}

func (c CartDetailRepositoryImpl) FindByCartIdAndProductId(cartId uuid.UUID, productId int) *domain.CartDetail {
	//TODO implement me
	panic("implement me")
}

func NewCartDetailRepository(db *gorm.DB) domain.CartDetailRepository {
	return &CartDetailRepositoryImpl{
		db: db,
	}
}

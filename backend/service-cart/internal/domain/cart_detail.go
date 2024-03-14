package domain

import "github.com/google/uuid"

type CartDetail struct {
	Id        int       `gorm:"id"`
	CartId    uuid.UUID `gorm:"cart_id"`
	ProductId int       `gorm:"product_id"`
	Qty       int       `gorm:"qty"`
}

type CartDetailRepository interface {
	InsertCartDetail(detail CartDetail) error
	UpdateCartDetail(detail CartDetail) error
	FindByCartIdAndProductId(cartId uuid.UUID, productId int) *CartDetail
}

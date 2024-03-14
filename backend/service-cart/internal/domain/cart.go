package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const (
	InOrder  = "IN_ORDER"
	InBasket = "IN_BASKET"
)

type Cart struct {
	Id          uuid.UUID `gorm:"id"`
	UserId      int       `gorm:"userId"`
	TotalAmount int       `gorm:"total_amount"`
	StatusCart  string    `gorm:"status_cart"`
	CreateAt    time.Time `gorm:"create_at"`
}

type CartRepository interface {
	CreateNewCart(cart Cart, tx gorm.DB) error
	FindCurrentCart(userId int, tx gorm.DB) *Cart
	UpdateCart(cart Cart, tx gorm.DB) error
}

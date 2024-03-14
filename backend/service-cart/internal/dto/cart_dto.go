package dto

import "service-cart/internal/domain"

type MyCartResponse struct {
	Cart       domain.Cart         `json:"cart"`
	DetailCart []domain.CartDetail `json:"detail_cart"`
}

type AddToCartRequest struct {
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
}

type AddToCartResponse struct {
	StatusToCart string `json:"status_to_cart"`
}

package dto

type ProductStockRequest struct {
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
}

type ProductStockResponse struct {
	Price          int  `json:"price"`
	IsStockReduced bool `json:"is_stock_reduced"`
}

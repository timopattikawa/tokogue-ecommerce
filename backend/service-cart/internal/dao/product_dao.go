package dao

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"service-cart/internal/dto"
)

type ProductDaoImpl struct {
	url string
}

func (p ProductDaoImpl) StockAvail(request dto.ProductStockRequest) *dto.ProductStockResponse {
	marshal, err := json.Marshal(request)
	if err != nil {
		log.Printf("fail to marshal, err: %s", err)
		return nil
	}

	reader := bytes.NewReader(marshal)
	resp, err := http.Post(p.url, "application/json", reader)
	if err != nil {
		log.Printf("Fail to send http post : %s", err)
		return nil
	}

	var responseBody = dto.ProductStockResponse{}

	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&responseBody)
	if err != nil {
		log.Printf("Fail to decode")
		return nil
	}

	if resp.StatusCode != 200 {
		log.Printf("Got Resp status %v", resp.StatusCode)
		return nil
	}

	return &responseBody
}

type ProductDao interface {
	StockAvail(request dto.ProductStockRequest) *dto.ProductStockResponse
}

func NewProductDao(url string) ProductDao {
	return &ProductDaoImpl{
		url: url,
	}
}

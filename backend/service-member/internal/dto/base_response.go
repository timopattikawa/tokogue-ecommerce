package dto

type BaseResponse[K comparable, T comparable] struct {
	StatusCode   int
	ErrorMessage K
	Data         T
}

package dto

type SendOtpKafka struct {
	Otp   string `json:"otp"`
	Email string `json:"email"`
}

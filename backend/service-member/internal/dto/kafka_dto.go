package dto

type SendOtpKafka struct {
	Type  string `json:"type"`
	Email string `json:"email"`
}

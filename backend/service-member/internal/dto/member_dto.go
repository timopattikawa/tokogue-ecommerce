package dto

import "time"

type MemberSignUpRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type MemberSignInRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type MemberEmailRequest struct {
	Email string `json:"email"`
}

type MemberAuthResponse struct {
	Token string `json:"token"`
}

type MemberResponse struct {
	Id       int       `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	CreateAt time.Time `json:"create_at" db:"create_at"`
}

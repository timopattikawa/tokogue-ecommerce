package domain

import (
	"service-member/internal/dto"
	"time"
)

type Member struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Password  string    `json:"password" db:"password"`
	AccessKey string    `json:"access_key" db:"access_key"`
	Email     string    `json:"email" db:"email"`
	CreateAt  time.Time `json:"create_at" db:"create_at"`
}

type MemberRepository interface {
	CreateMember(member Member) error
	FindMemberById(id int) (Member, error)
	FindMemberByEmail(email string) (Member, error)
	DeleteMemberById(id Member) (Member, error)
}

type MemberService interface {
	SignUpMember(request dto.MemberSignUpRequest) (dto.MemberAuthResponse, error)
	SignInMember(request dto.MemberSignInRequest) (dto.MemberAuthResponse, error)
	DeleteMember(header dto.Header) (string, error)
	GetMemberById(id int) (dto.MemberResponse, error)
}

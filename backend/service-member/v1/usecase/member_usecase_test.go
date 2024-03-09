package usecase

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"service-member/internal/domain"
	"service-member/internal/dto"
	"service-member/internal/exception"
	mock2 "service-member/internal/util/mock"
	"service-member/v1/repository"
	"testing"
	"time"
)

var now = time.Now

func TestGetMemberById_FailToGet404(t *testing.T) {
	var mockRepository = repository.MemberRepositoryMock{Mock: mock.Mock{}}
	var mockJwtGenerator = mock2.JwtGeneratorImplMock{Mock: mock.Mock{}}
	var memberUsecase = MemberUsecaseImpl{memberRepo: &mockRepository, jwtGenerator: &mockJwtGenerator}

	mockRepository.Mock.On("FindMemberById", 101).Return(nil, fmt.Errorf(""))

	result, err := memberUsecase.GetMemberById(101)
	assert.Equal(t, result, dto.MemberResponse{})
	assert.Equal(t, err, exception.NotFoundError{Message: "Member not found"})
}

func TestGetMemberById_SuccessToGetMember(t *testing.T) {
	var mockRepository = repository.MemberRepositoryMock{Mock: mock.Mock{}}
	var mockJwtGenerator = mock2.JwtGeneratorImplMock{Mock: mock.Mock{}}
	var memberUsecase = MemberUsecaseImpl{memberRepo: &mockRepository, jwtGenerator: &mockJwtGenerator}

	var memberDummy = domain.Member{
		Id:        101,
		Name:      "Timo Pattikawa",
		Password:  "asdfasdf",
		AccessKey: "TMPTMPTMPTMP",
		CreateAt:  time.Now(),
	}
	mockRepository.Mock.On("FindMemberById", 101).Return(memberDummy)

	result, err := memberUsecase.GetMemberById(101)
	assert.NotNil(t, result)
	assert.Equal(t, memberDummy.Id, result.Id)
	assert.Equal(t, memberDummy.Name, result.Name)
	assert.Nil(t, err)

}

func TestSignUp_SuccessToSignUp(t *testing.T) {
	var mockRepository = repository.MemberRepositoryMock{Mock: mock.Mock{}}
	var mockJwtGenerator = mock2.JwtGeneratorImplMock{Mock: mock.Mock{}}
	var mockPassword = mock2.PasswordGeneratorMock{Mock: mock.Mock{}}

	var memberUsecase = MemberUsecaseImpl{
		memberRepo:   &mockRepository,
		jwtGenerator: &mockJwtGenerator,
		hashPassword: &mockPassword,
	}

	request := dto.MemberSignUpRequest{
		Name:     "Timo Test",
		Password: "asdfasdf",
		Email:    "timo@gmail.com",
	}

	member := domain.Member{
		Name:      request.Name,
		Password:  "asdfasdf",
		AccessKey: "asdfasdf",
		Email:     request.Email,
	}

	mockRepository.Mock.On("FindMemberByEmail", request.Email).Return(nil, fmt.Errorf(""))
	mockRepository.Mock.On("CreateMember", member).Return(nil)
	mockJwtGenerator.Mock.On("NewAccessToken", request.Email, request.Name).Return("asdfasdf", nil)
	mockPassword.Mock.On("HashPassword", request.Password).Return("asdfasdf", nil)

	result, err := memberUsecase.SignUpMember(request)
	assert.NotNil(t, result)
	assert.Equal(t, result, dto.MemberAuthResponse{Token: "asdfasdf"})
	assert.Nil(t, err)
}

func TestSignUp_FailBecauseMemberHasCreated(t *testing.T) {
	var mockRepository = repository.MemberRepositoryMock{Mock: mock.Mock{}}
	var mockJwtGenerator = mock2.JwtGeneratorImplMock{Mock: mock.Mock{}}
	var mockPassword = mock2.PasswordGeneratorMock{Mock: mock.Mock{}}

	var memberUsecase = MemberUsecaseImpl{
		memberRepo:   &mockRepository,
		jwtGenerator: &mockJwtGenerator,
		hashPassword: &mockPassword,
	}
	request := dto.MemberSignUpRequest{
		Name:     "Timo Test",
		Password: "asdfasdf",
		Email:    "timo@gmail.com",
	}

	member := domain.Member{
		Name:      request.Name,
		Password:  request.Password,
		AccessKey: "asdfasdf",
		Email:     request.Email,
	}

	mockRepository.Mock.On("FindMemberByEmail", request.Email).Return(member)

	result, err := memberUsecase.SignUpMember(request)
	assert.Equal(t, result, dto.MemberAuthResponse{})
	assert.Equal(t, err, exception.BadRequestError{Message: "Member has been created"})
}

func TestSignUp_FailBecauseFailHashPassword(t *testing.T) {
	var mockRepository = repository.MemberRepositoryMock{Mock: mock.Mock{}}
	var mockJwtGenerator = mock2.JwtGeneratorImplMock{Mock: mock.Mock{}}
	var mockPassword = mock2.PasswordGeneratorMock{Mock: mock.Mock{}}

	var memberUsecase = MemberUsecaseImpl{
		memberRepo:   &mockRepository,
		jwtGenerator: &mockJwtGenerator,
		hashPassword: &mockPassword,
	}
	request := dto.MemberSignUpRequest{
		Name:     "Timo Test",
		Password: "asdfasdf",
		Email:    "timo@gmail.com",
	}

	mockRepository.Mock.On("FindMemberByEmail", request.Email).Return(nil, fmt.Errorf(""))
	mockJwtGenerator.Mock.On("NewAccessToken", request.Email, request.Name).Return("",
		fmt.Errorf("Error jwt generate"))

	result, err := memberUsecase.SignUpMember(request)
	assert.Equal(t, result, dto.MemberAuthResponse{})
	assert.NotNil(t, err)
}

func TestSignUp_FailBecauseFailJWT(t *testing.T) {
	var mockRepository = repository.MemberRepositoryMock{Mock: mock.Mock{}}
	var mockJwtGenerator = mock2.JwtGeneratorImplMock{Mock: mock.Mock{}}
	var mockPassword = mock2.PasswordGeneratorMock{Mock: mock.Mock{}}

	var memberUsecase = MemberUsecaseImpl{
		memberRepo:   &mockRepository,
		jwtGenerator: &mockJwtGenerator,
		hashPassword: &mockPassword,
	}
	request := dto.MemberSignUpRequest{
		Name:     "Timo Test",
		Password: "asdfasdf",
		Email:    "timo@gmail.com",
	}

	mockRepository.Mock.On("FindMemberByEmail", request.Email).Return(nil, fmt.Errorf(""))
	mockJwtGenerator.Mock.On("NewAccessToken", request.Email, request.Name).Return("", nil)
	mockPassword.Mock.On("HashPassword", request.Password).Return("", fmt.Errorf("Err"))

	result, err := memberUsecase.SignUpMember(request)
	assert.Equal(t, result, dto.MemberAuthResponse{})
	assert.NotNil(t, err)
}

func TestSignIn_SuccessToSignIn(t *testing.T) {
	var mockRepository = repository.MemberRepositoryMock{Mock: mock.Mock{}}
	var mockJwtGenerator = mock2.JwtGeneratorImplMock{Mock: mock.Mock{}}
	var mockPassword = mock2.PasswordGeneratorMock{Mock: mock.Mock{}}

	var memberUsecase = MemberUsecaseImpl{
		memberRepo:   &mockRepository,
		jwtGenerator: &mockJwtGenerator,
		hashPassword: &mockPassword,
	}

	request := dto.MemberSignInRequest{
		Password: "asdfasdf",
		Email:    "timo@gmail.com",
	}
	member := domain.Member{
		Name:      "Timo Test",
		Password:  request.Password,
		AccessKey: "token",
		Email:     request.Email,
	}

	mockRepository.Mock.On("FindMemberByEmail", request.Email).Return(member, nil)
	mockPassword.Mock.On("CheckPasswordHash", request.Password, member.Password).Return(true)
	result, err := memberUsecase.SignInMember(request)

	assert.Equal(t, result.Token, member.AccessKey)
	assert.Nil(t, err)
}

func TestSignIn_FailToSignInMember404(t *testing.T) {
	var mockRepository = repository.MemberRepositoryMock{Mock: mock.Mock{}}
	var mockJwtGenerator = mock2.JwtGeneratorImplMock{Mock: mock.Mock{}}
	var mockPassword = mock2.PasswordGeneratorMock{Mock: mock.Mock{}}

	var memberUsecase = MemberUsecaseImpl{
		memberRepo:   &mockRepository,
		jwtGenerator: &mockJwtGenerator,
		hashPassword: &mockPassword,
	}

	request := dto.MemberSignInRequest{
		Password: "asdfasdf",
		Email:    "timo@gmail.com",
	}
	member := domain.Member{
		Name:      "Timo Test",
		Password:  request.Password,
		AccessKey: "token",
		Email:     request.Email,
	}

	mockRepository.Mock.On("FindMemberByEmail", request.Email).Return(nil, fmt.Errorf(""))
	mockPassword.Mock.On("CheckPasswordHash", request.Password, member.Password).Return(true)
	result, err := memberUsecase.SignInMember(request)

	assert.Equal(t, result, dto.MemberAuthResponse{})
	assert.Equal(t, err, exception.NotFoundError{Message: "Member not found"})
}

func TestSignIn_FailToSignInPasswordNotMatch(t *testing.T) {
	var mockRepository = repository.MemberRepositoryMock{Mock: mock.Mock{}}
	var mockJwtGenerator = mock2.JwtGeneratorImplMock{Mock: mock.Mock{}}
	var mockPassword = mock2.PasswordGeneratorMock{Mock: mock.Mock{}}

	var memberUsecase = MemberUsecaseImpl{
		memberRepo:   &mockRepository,
		jwtGenerator: &mockJwtGenerator,
		hashPassword: &mockPassword,
	}

	request := dto.MemberSignInRequest{
		Password: "asdfasdf",
		Email:    "timo@gmail.com",
	}
	member := domain.Member{
		Name:      "Timo Test",
		Password:  "asdf",
		AccessKey: "token",
		Email:     request.Email,
	}

	mockRepository.Mock.On("FindMemberByEmail", request.Email).Return(member, nil)
	mockPassword.Mock.On("CheckPasswordHash", request.Password, member.Password).Return(false)
	result, err := memberUsecase.SignInMember(request)

	assert.Equal(t, result, dto.MemberAuthResponse{})
	assert.Equal(t, err, exception.BadRequestError{Message: "Invalid Password please check again!!"})
}

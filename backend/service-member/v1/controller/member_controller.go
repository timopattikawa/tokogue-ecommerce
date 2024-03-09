package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"service-member/internal/domain"
	"service-member/internal/dto"
	"service-member/internal/exception"
	"strconv"
)

type MemberController struct {
	memberUsecase domain.MemberService
}

func NewMemberController(service domain.MemberService) MemberController {
	return MemberController{
		memberUsecase: service,
	}
}

func (m *MemberController) SignUpMember(c *fiber.Ctx) error {
	body := c.Body()

	var requestSignUp = dto.MemberSignUpRequest{}
	err := json.Unmarshal(body, &requestSignUp)
	if err != nil {
		log.Printf("fail unmarshal body err: {%s}", err)
		return err
	}

	member, err := m.memberUsecase.SignUpMember(requestSignUp)
	if err != nil {
		log.Printf("Error when signup err: {%s}", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.BaseResponse[interface{}, dto.MemberAuthResponse]{
		StatusCode:   fiber.StatusOK,
		ErrorMessage: nil,
		Data:         member,
	})
}

func (m *MemberController) SignInMember(c *fiber.Ctx) error {
	body := c.Body()

	var requestSignIn = dto.MemberSignInRequest{}
	err := json.Unmarshal(body, &requestSignIn)
	if err != nil {
		log.Printf("Fail to unmarshal body err {%s}", err.Error())
		return err
	}

	member, err := m.memberUsecase.SignInMember(requestSignIn)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.BaseResponse[interface{}, dto.MemberAuthResponse]{
		StatusCode:   fiber.StatusOK,
		ErrorMessage: nil,
		Data:         member,
	})
}

func (m *MemberController) GetMemberByEmail(c *fiber.Ctx) error {
	params := c.Query("id")

	memberRequest, err := strconv.Atoi(params)
	if err != nil {
		log.Printf("Err : {%s}", err.Error())
		return exception.BadRequestError{Message: "Bad request please check your body or params"}
	}

	memberById, err := m.memberUsecase.GetMemberById(memberRequest)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.BaseResponse[interface{}, dto.MemberResponse]{
		StatusCode:   fiber.StatusOK,
		ErrorMessage: nil,
		Data:         memberById,
	})
}

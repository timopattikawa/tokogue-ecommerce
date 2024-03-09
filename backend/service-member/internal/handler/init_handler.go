package handler

import (
	"github.com/gofiber/fiber/v2"
	"service-member/v1/controller"
)

func NewHandlerMember(app *fiber.App, handler controller.MemberController) {
	app.Post("/v1/signup", handler.SignUpMember)
	app.Post("/v1/signin", handler.SignInMember)
	app.Get("/v1/get/member", handler.GetMemberByEmail)
}

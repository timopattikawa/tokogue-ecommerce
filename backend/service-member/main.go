package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"service-member/config"
	"service-member/internal/db"
	"service-member/internal/exception"
	"service-member/internal/handler"
	"service-member/internal/util"
	"service-member/v1/controller"
	"service-member/v1/repository"
	"service-member/v1/usecase"
)

func main() {
	cfg := config.NewConfig()

	postgres := db.NewConnectionPostgres(cfg)
	memberRepository := repository.NewMemberRepository(postgres)
	jwtGenerator := util.NewJWTGenerator()
	passwordGenerator := util.NewPasswordGenerator()
	kafkaConfig := util.NewKafkaConfig("localhost:9092", "tokoguemessage", 3)
	memberUsecase := usecase.NewUsecaseMember(memberRepository, jwtGenerator, passwordGenerator, kafkaConfig)
	memberController := controller.NewMemberController(memberUsecase)

	app := fiber.New(
		fiber.Config{
			ErrorHandler: exception.ErrorHandler,
		})

	handler.NewHandlerMember(app, memberController)

	err := app.Listen(":9001")
	if err != nil {
		log.Fatal("Fail To Listen 9001")
	}
}

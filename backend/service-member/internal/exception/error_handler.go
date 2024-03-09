package exception

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"service-member/internal/dto"
	"time"
)

func ErrorHandler(c *fiber.Ctx, err error) error {

	var notFoundError NotFoundError
	errNotFound := errors.As(err, &notFoundError)
	if errNotFound {
		return c.Status(fiber.StatusNotFound).JSON(dto.BaseResponse[ErrorResponse, interface{}]{
			StatusCode: fiber.StatusNotFound,
			ErrorMessage: ErrorResponse{
				Message: notFoundError.Message,
				Date:    time.Now(),
			},
		})
	}

	var badRequestError BadRequestError
	errBadReq := errors.As(err, &badRequestError)
	if errBadReq {
		return c.Status(fiber.StatusBadRequest).JSON(dto.BaseResponse[ErrorResponse, interface{}]{
			StatusCode: fiber.StatusBadRequest,
			ErrorMessage: ErrorResponse{
				Message: badRequestError.Message,
				Date:    time.Now(),
			},
		})
	}

	log.Printf("Error Catch : {%s}", err.Error())
	return c.Status(fiber.StatusInternalServerError).JSON(dto.BaseResponse[ErrorResponse, interface{}]{
		StatusCode: fiber.StatusInternalServerError,
		ErrorMessage: ErrorResponse{
			Message: "Something wrong with our server, please try again!!",
			Date:    time.Now(),
		},
	})
}

package exception

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"time"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var notFoundError NotFoundError

	if as := errors.As(err, &notFoundError); as {
		return c.Status(fiber.StatusNotFound).JSON(BaseError{
			status:  fiber.StatusNotFound,
			message: err.Error(),
			date:    time.Now(),
		})
	}

	var badRequest BadRequestError
	if as := errors.As(err, &badRequest); as {
		return c.Status(fiber.StatusBadRequest).JSON(
			BaseError{
				status:  fiber.StatusBadRequest,
				message: err.Error(),
				date:    time.Now(),
			})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(
		BaseError{
			status:  fiber.StatusInternalServerError,
			message: err.Error(),
			date:    time.Now(),
		})
}

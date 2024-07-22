package error_handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gogram/pkg/helpers"
)

func GlobalErrorHandler(ctx *fiber.Ctx, err error) error {
	res := helpers.NewResponseHelper(
		fiber.StatusInternalServerError,
		err.Error(),
		nil,
		nil,
	)

	var e *fiber.Error

	if errors.As(err, &e) {
		res.Code = e.Code
	}

	return ctx.Status(res.Code).JSON(res)
}

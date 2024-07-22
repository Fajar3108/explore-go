package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	if err := validate.RegisterValidation("image", imageValidation); err != nil {
		panic(err)
	}
}

func Validator() *validator.Validate {
	return validate
}

func FiberValidationError(err error) *fiber.Error {
	for _, err := range err.(validator.ValidationErrors) {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return fiber.NewError(fiber.StatusBadRequest, "Bad request")
}

func Validate[T any](ctx *fiber.Ctx, request *T) (err error) {
	err = ctx.BodyParser(request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if errs := Validator().Struct(request); errs != nil {
		return FiberValidationError(errs)
	}

	return nil
}

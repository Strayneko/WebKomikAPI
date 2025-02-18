package handlers

import (
	"errors"
	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ErrorHandler struct {
	errorHandler interfaces.HandlerInterface
}

func NewErrorHandler() interfaces.HandlerInterface {
	return &ErrorHandler{}
}

func (h *ErrorHandler) ErrorHandler(ctx *fiber.Ctx, err error) error {
	var e *fiber.Error
	code := http.StatusOK

	if errors.As(err, &e) {
		code = e.Code
	}

	return ctx.Status(code).JSON(&types.ResponseType{
		Status:  false,
		Code:    int16(code),
		Message: err.Error(),
	})
}

package render

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func SendError(ctx *fiber.Ctx, statusCode int, err error, message string) error {

	return ctx.Status(statusCode).JSON(ResponseError{
		Error:      err.Error(),
		Message:    message,
		StatusCode: statusCode,
	})

}

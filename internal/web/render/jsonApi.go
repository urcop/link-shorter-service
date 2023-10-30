package render

import "github.com/gofiber/fiber/v2"

func JSONAPIPayload(ctx *fiber.Ctx, statusCode int, message string) error {

	return ctx.Status(statusCode).JSON(ResponseError{
		Message:    message,
		StatusCode: statusCode,
	})

}

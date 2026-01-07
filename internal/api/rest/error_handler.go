package rest

import "github.com/gofiber/fiber/v3"

func errorHandler(fCtx fiber.Ctx, err error) error {
	// todo:improve with dynamic status
	const status = fiber.StatusInternalServerError

	fCtx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	fCtx.Status(status).JSON(fiber.Map{
		"message": err.Error(),
	})

	return nil
}

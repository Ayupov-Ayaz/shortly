package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"github.com/ayupov-ayaz/shortly/internal/entity"
	"github.com/ayupov-ayaz/shortly/internal/service/shorten"
)

type URLHandler struct {
	srv shorten.Shortener
}

func NewURLHandler(srv shorten.Shortener) *URLHandler {
	return &URLHandler{
		srv: srv,
	}
}

func (h *URLHandler) RegisterRouter(router fiber.Router) {
	router.Post(shortenPath, h.shortenURL)
}

func (h *URLHandler) shortenURL(fCtx fiber.Ctx) error {
	req := new(entity.CreateURLRequest)

	err := fCtx.Bind().JSON(req)
	if err != nil {
		//todo: log
		return fmt.Errorf("bind req body: %w", err)
	}

	err = req.Validate()
	if err != nil {
		//todo: log
		return fmt.Errorf("validate req: %w", err)
	}

	resp, err := h.srv.CreateShortURL(fCtx.Context(), req)
	if err != nil {
		// todo: log
		return fmt.Errorf("creating short url: %w", err)
	}

	fCtx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	err = fCtx.Status(fiber.StatusOK).JSON(resp)
	if err != nil {
		// todo: log
		return fmt.Errorf("sending json: %w", err)
	}

	return nil
}

package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ayupov-ayaz/shortly/internal/api/gen"
	"github.com/ayupov-ayaz/shortly/internal/service/shortener"
)

type URLShortener struct {
	srv        shortener.Shortener
	respWriter *ResponseWriter
}

func NewURLShortener(
	srv shortener.Shortener,
	respWriter *ResponseWriter,
) *URLShortener {
	return &URLShortener{
		srv:        srv,
		respWriter: respWriter,
	}
}

func (h *URLShortener) CreateShortURL(
	writer http.ResponseWriter, httpReq *http.Request,
) {
	var req gen.CreateURLRequest

	err := json.NewDecoder(httpReq.Body).Decode(&req)
	if err != nil {
		log.Printf("decode request: %v", err) // todo: zerolog
		h.respWriter.SendBadRequest(writer, errors.New("invalid request"))
		return
	}

	resp, err := h.srv.ShortenURL(httpReq.Context(), req)
	if err != nil {
		log.Printf("shorten url: %v", err) // todo: zerolog
		h.respWriter.SendError(writer, err)
		return
	}

	h.respWriter.SendCreated(writer, resp)
}

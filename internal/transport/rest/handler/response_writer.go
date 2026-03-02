package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ayupov-ayaz/shortly/internal/api/gen"
)

type ResponseWriter struct{}

func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{}
}

func (r *ResponseWriter) SendCreated(
	writer http.ResponseWriter, resp any,
) {
	r.sendJson(writer, http.StatusCreated, resp)
}

func (r *ResponseWriter) SendOk(
	writer http.ResponseWriter, resp any,
) {
	r.sendJson(writer, http.StatusOK, resp)
}

func (r *ResponseWriter) SendError(
	writer http.ResponseWriter, resp error,
) {
	// todo: check is redis error => internal server error
	// todo: check is postgres error => internal server error
	r.SendBadRequest(writer, resp)
}

func (r *ResponseWriter) SendInternalServerError(
	writer http.ResponseWriter, resp error,
) {
	r.sendError(writer, http.StatusInternalServerError, resp)
}

func (r *ResponseWriter) SendBadRequest(
	writer http.ResponseWriter, resp error,
) {
	r.sendError(writer, http.StatusBadRequest, resp)
}

func (r *ResponseWriter) sendError(
	writer http.ResponseWriter, status int, err error,
) {
	resp := gen.Error{
		Error: err.Error(),
	}

	r.sendJson(writer, status, resp)
}

func (r *ResponseWriter) sendJson(
	writer http.ResponseWriter, status int, data any,
) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	if err := json.NewEncoder(writer).Encode(data); err != nil {
		// todo: zerolog
		log.Printf("encode response: %v", err)
	}
}

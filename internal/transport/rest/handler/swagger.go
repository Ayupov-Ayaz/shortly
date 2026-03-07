package handler

import (
	"net/http"

	"github.com/ayupov-ayaz/shortly/internal/api/gen"
)

type Swagger struct {
	respWriter *ResponseWriter
}

func NewSwagger(respWriter *ResponseWriter) *Swagger {
	return &Swagger{
		respWriter: respWriter,
	}
}

func (h *Swagger) GetSwaggerUI(
	writer http.ResponseWriter, httpReq *http.Request,
) {
	swagger, err := gen.GetSwagger()
	if err != nil {
		h.respWriter.SendInternalServerError(writer, err)
		return
	}

	data, err := swagger.MarshalJSON()
	if err != nil {
		h.respWriter.SendInternalServerError(writer, err)
		return
	}

	_, err = writer.Write(data)
	if err != nil {
		h.respWriter.SendInternalServerError(writer, err)
	}
}
